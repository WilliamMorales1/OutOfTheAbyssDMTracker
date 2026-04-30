package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

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
	var sb strings.Builder
	sb.WriteByte('[')
	for i, f := range er.Embedding {
		if i > 0 {
			sb.WriteByte(',')
		}
		fmt.Fprintf(&sb, "%g", f)
	}
	sb.WriteByte(']')
	return sb.String(), nil
}

type searchResult struct {
	ChapterTitle string
	Content      string
	Score        float64
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<p class="text-secondary">Enter a query above.</p>`)
		return
	}

	emb, err := queryEmbedding(r.Context(), q)
	if err != nil {
		http.Error(w, "Embedding error: "+err.Error(), 500)
		return
	}

	rows, err := db.Query(r.Context(), `
		SELECT chapter_title, content,
		       1 - (embedding <=> $1::vector) AS score
		FROM chapter_chunks
		ORDER BY embedding <=> $1::vector
		LIMIT 5`, emb)
	if err != nil {
		http.Error(w, "Search error: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	var results []searchResult
	for rows.Next() {
		var sr searchResult
		if err := rows.Scan(&sr.ChapterTitle, &sr.Content, &sr.Score); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		results = append(results, sr)
	}

	funcs := template.FuncMap{
		"pct": func(f float64) string { return fmt.Sprintf("%.0f", f*100) },
	}
	tmpl := template.Must(template.New("r").Funcs(funcs).Parse(searchResultsTmpl))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, results)
}

var searchResultsTmpl = `
{{if .}}
  {{range .}}
  <div class="card bg-dark border-secondary mb-3">
    <div class="card-header d-flex justify-content-between align-items-center">
      <strong class="text-warning">{{.ChapterTitle}}</strong>
      <span class="badge bg-secondary">{{pct .Score}}% match</span>
    </div>
    <div class="card-body text-light" style="white-space:pre-wrap;font-size:.875rem">{{.Content}}</div>
  </div>
  {{end}}
{{else}}
  <p class="text-secondary">No results found.</p>
{{end}}`
