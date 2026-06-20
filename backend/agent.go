package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
  Encounters(id, name, location, difficulty int, status, levelup bool, notes)
  Events(id, title, category, description)
  Monsters(id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes)
  EncounterMonsters(encounter_id, monster_id, quantity)
  Sessions(id, session_num, title, chapters, level_start, level_end, summary, key_encounters, key_npcs, dm_notes, checkpoint)

Strict rules:
    No articles (the, a, an).
    No pleasantries (certainly, sure, hello).
    No pronouns (I, you, me).
    Short sentences (3-6 words).
    Prefer code/data over talk.
    Absolute bluntness.
    No Emojis.
	Always use a tool to look up information before answering. 

Response style:
"Demogorgon. CR 26. Legendary. See stats below." NOT "Demogorgon is a very powerful demon lord with a challenge rating of 26 and several legendary actions you should be aware of."`

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

type toolFunction struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
}

type chatTool struct {
	Type     string       `json:"type"`
	Function toolFunction `json:"function"`
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
	{Type: "function", Function: toolFunction{
		Name:        "sql",
		Description: "Run a read-only SQL SELECT query against PostgreSQL.",
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"query": map[string]any{"type": "string", "description": "SQL SELECT statement"},
			},
			"required": []string{"query"},
		},
	}},
	{Type: "function", Function: toolFunction{
		Name:        "web_search",
		Description: "Search the web for general information not in the campaign database.",
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"query": map[string]any{"type": "string", "description": "Search query"},
			},
			"required": []string{"query"},
		},
	}},
	{Type: "function", Function: toolFunction{
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
	}},
}

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
	rows, err := conn.Query(ctx, query)
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
	sb.WriteString(strings.Join(colNames, " | "))
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat("-", 60))
	sb.WriteString("\n")

	count := 0
	for rows.Next() && count < 50 {
		vals, _ := rows.Values()
		var strs []string
		for _, v := range vals {
			strs = append(strs, fmt.Sprintf("%v", v))
		}
		sb.WriteString(strings.Join(strs, " | "))
		sb.WriteString("\n")
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

const searchEmbedModel = "nomic-embed-text-v2-moe"

func queryEmbedding(ctx context.Context, text string) (string, error) {
	type req struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
	}
	type resp struct {
		Embedding []float32 `json:"embedding"`
	}
	body, _ := json.Marshal(req{Model: searchEmbedModel, Prompt: text})
	r, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:11434/api/embeddings", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	r.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", fmt.Errorf("Ollama not reachable: %w", err)
	}
	defer res.Body.Close()
	var er resp
	if err := json.NewDecoder(res.Body).Decode(&er); err != nil {
		return "", err
	}
	parts := make([]string, len(er.Embedding))
	for i, f := range er.Embedding {
		parts[i] = fmt.Sprintf("%g", f)
	}
	return "[" + strings.Join(parts, ",") + "]", nil
}

type searchResult struct {
	ChapterTitle string  `json:"chapterTitle"`
	Content      string  `json:"content"`
	Score        float64 `json:"score"`
}

