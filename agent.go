package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const ollamaURL = "http://localhost:11434/v1/chat/completions"
const agentModel = "gemma4"

const systemPrompt = `You are a D&D Dungeon Master assistant for "Out of the Abyss".
Use the search and sql tools to answer questions about the campaign. Use web_search for general knowledge not in the database.

Database schema:
  Locations(id, name, type, status, danger int 1-5, description, secrets)
  NPCS(id, name, madness int, disposition, location, notes, description)
  Encounters(id, name, location, difficulty int, status, enemies, levelup bool, notes)
  Events(id, title, category, description)

Always use a tool to look up information before answering. Be concise.`

type chatMsg struct {
	Role       string     `json:"role"`
	Content    string     `json:"content,omitempty"`
	ToolCalls  []toolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
}

type toolCall struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

type chatTool struct {
	Type     string `json:"type"`
	Function struct {
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Parameters  map[string]any `json:"parameters"`
	} `json:"function"`
}

type chatReq struct {
	Model    string     `json:"model"`
	Messages []chatMsg  `json:"messages"`
	Tools    []chatTool `json:"tools,omitempty"`
	Stream   bool       `json:"stream"`
}

type chatResp struct {
	Choices []struct {
		Message      chatMsg `json:"message"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
}

var tools = []chatTool{
	{
		Type: "function",
		Function: struct {
			Name        string         `json:"name"`
			Description string         `json:"description"`
			Parameters  map[string]any `json:"parameters"`
		}{
			Name:        "search",
			Description: "Full-text fuzzy search across the campaign database.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query": map[string]any{"type": "string", "description": "Search text"},
					"table": map[string]any{
						"type":        "string",
						"description": "Table: locations, npcs, encounters, events, or all",
						"enum":        []string{"locations", "npcs", "encounters", "events", "all"},
					},
				},
				"required": []string{"query", "table"},
			},
		},
	},
	{
		Type: "function",
		Function: struct {
			Name        string         `json:"name"`
			Description string         `json:"description"`
			Parameters  map[string]any `json:"parameters"`
		}{
			Name:        "sql",
			Description: "Run a read-only SQL SELECT query against PostgreSQL.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query": map[string]any{"type": "string", "description": "SQL SELECT statement"},
				},
				"required": []string{"query"},
			},
		},
	},
	{
		Type: "function",
		Function: struct {
			Name        string         `json:"name"`
			Description string         `json:"description"`
			Parameters  map[string]any `json:"parameters"`
		}{
			Name:        "web_search",
			Description: "Search the web for general information not in the campaign database.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query": map[string]any{"type": "string", "description": "Search query"},
				},
				"required": []string{"query"},
			},
		},
	},
	{
		Type: "function",
		Function: struct {
			Name        string         `json:"name"`
			Description string         `json:"description"`
			Parameters  map[string]any `json:"parameters"`
		}{
			Name:        "dnd_lookup",
			Description: "Look up D&D 5e rules data (monsters, spells, equipment, classes, etc.) from the official 5e API.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"category": map[string]any{
						"type":        "string",
						"description": "Resource category, e.g. monsters, spells, equipment, magic-items, classes, races, conditions, damage-types",
					},
					"index": map[string]any{
						"type":        "string",
						"description": "Specific resource slug (e.g. 'aboleth', 'fireball'). Omit to list all in category.",
					},
				},
				"required": []string{"category"},
			},
		},
	},
}

func initAI() {}

func ollama(ctx context.Context, messages []chatMsg) (*chatResp, error) {
	body, _ := json.Marshal(chatReq{
		Model:    agentModel,
		Messages: messages,
		Tools:    tools,
		Stream:   false,
	})
	req, err := http.NewRequestWithContext(ctx, "POST", ollamaURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Ollama not reachable - is it running? (%w)", err)
	}
	defer resp.Body.Close()
	var cr chatResp
	json.NewDecoder(resp.Body).Decode(&cr)
	return &cr, nil
}

func runAgent(ctx context.Context, question string) (string, error) {
	messages := []chatMsg{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: question},
	}

	for range 8 {
		resp, err := ollama(ctx, messages)
		if err != nil {
			return "", err
		}
		if len(resp.Choices) == 0 {
			return "No response from model.", nil
		}

		msg := resp.Choices[0].Message
		messages = append(messages, msg)

		if resp.Choices[0].FinishReason != "tool_calls" || len(msg.ToolCalls) == 0 {
			return msg.Content, nil
		}

		for _, tc := range msg.ToolCalls {
			result := executeTool(ctx, tc.Function.Name, tc.Function.Arguments)
			messages = append(messages, chatMsg{
				Role:       "tool",
				ToolCallID: tc.ID,
				Content:    result,
			})
		}
	}
	return "Max tool iterations reached.", nil
}

func executeTool(ctx context.Context, name, argsJSON string) string {
	var args map[string]any
	if err := json.Unmarshal([]byte(argsJSON), &args); err != nil {
		return "Error parsing args: " + err.Error()
	}
	switch name {
	case "search":
		query, _ := args["query"].(string)
		table, _ := args["table"].(string)
		return searchBleve(query, table)
	case "sql":
		query, _ := args["query"].(string)
		return execSQL(ctx, query)
	case "web_search":
		query, _ := args["query"].(string)
		return webSearch(ctx, query)
	case "dnd_lookup":
		category, _ := args["category"].(string)
		index, _ := args["index"].(string)
		return dndLookup(ctx, category, index)
	}
	return "Unknown tool: " + name
}

var reResult = regexp.MustCompile(`(?s)class="result__title"[^>]*>.*?<a[^>]*href="([^"]*)"[^>]*>(.*?)</a>.*?class="result__snippet"[^>]*>(.*?)</span>`)
var reTag = regexp.MustCompile(`<[^>]+>`)

func webSearch(ctx context.Context, query string) string {
	req, err := http.NewRequestWithContext(ctx, "GET",
		"https://html.duckduckgo.com/html/?q="+url.QueryEscape(query), nil)
	if err != nil {
		return "Error: " + err.Error()
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Error: " + err.Error()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	matches := reResult.FindAllSubmatch(body, 5)
	if len(matches) == 0 {
		return "No results found."
	}
	var sb strings.Builder
	for _, m := range matches {
		title := reTag.ReplaceAllString(string(m[2]), "")
		snippet := reTag.ReplaceAllString(string(m[3]), "")
		link := string(m[1])
		fmt.Fprintf(&sb, "- %s\n  %s\n  %s\n\n", strings.TrimSpace(title), strings.TrimSpace(snippet), link)
	}
	return sb.String()
}

const dndAPIBase = "https://www.dnd5eapi.co/api/2014"

func dndLookup(ctx context.Context, category, index string) string {
	url := dndAPIBase + "/" + category
	if index != "" {
		url += "/" + index
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "Error building request: " + err.Error()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Error calling D&D API: " + err.Error()
	}
	defer resp.Body.Close()
	var result any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "Error decoding response: " + err.Error()
	}
	out, _ := json.MarshalIndent(result, "", "  ")
	s := string(out)
	if len(s) > 4000 {
		s = s[:4000] + "\n...(truncated)"
	}
	return s
}

func execSQL(ctx context.Context, query string) string {
	q := strings.TrimSpace(strings.ToUpper(query))
	if !strings.HasPrefix(q, "SELECT") {
		return "Error: only SELECT queries allowed."
	}
	rows, err := db.Query(ctx, query)
	if err != nil {
		return "Query error: " + err.Error()
	}
	defer rows.Close()

	cols := rows.FieldDescriptions()
	var colNames []string
	for _, c := range cols {
		colNames = append(colNames, string(c.Name))
	}

	var sb strings.Builder
	sb.WriteString(strings.Join(colNames, " | ") + "\n")
	sb.WriteString(strings.Repeat("-", 60) + "\n")

	count := 0
	for rows.Next() && count < 50 {
		vals, _ := rows.Values()
		var strs []string
		for _, v := range vals {
			strs = append(strs, fmt.Sprintf("%v", v))
		}
		sb.WriteString(strings.Join(strs, " | ") + "\n")
		count++
	}
	if count == 0 {
		return "No rows returned."
	}
	if count == 50 {
		sb.WriteString("(truncated at 50 rows)\n")
	}
	return sb.String()
}

func handleChat(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		http.Error(w, "missing q", 400)
		return
	}
	answer, err := runAgent(r.Context(), q)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	answer = strings.ReplaceAll(answer, "\n\n", "</p><p>")
	answer = strings.ReplaceAll(answer, "\n", "<br>")
	fmt.Fprintf(w,
		`<div class="chat-msg user d-flex justify-content-end"><div class="chat-bubble w-1000">%s</div></div>`+
			`<div class="chat-msg agent d-flex justify-content-start"><div class="chat-bubble w-1000"><p>%s</p></div></div>`,
		template.HTMLEscapeString(q), answer)
}

func handleReindex(w http.ResponseWriter, r *http.Request) {
	if err := indexAll(); err != nil {
		http.Error(w, "Reindex failed: "+err.Error(), 500)
		return
	}
	fmt.Fprint(w, "Indexed ✓")
}
