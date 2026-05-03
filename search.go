package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	parts := make([]string, len(er.Embedding))
	for i, f := range er.Embedding {
		parts[i] = fmt.Sprintf("%g", f)
	}
	return "[" + strings.Join(parts, ",") + "]", nil
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

	rows, err := conn.Query(r.Context(), `
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

	w.Header().Set("Content-Type", "text/html")
	SearchResults(results).Render(r.Context(), w)
}
