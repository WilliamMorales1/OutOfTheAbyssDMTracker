// Package embeddings talks to a local Ollama instance to turn text into
// vector embeddings for semantic search.
package embeddings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const Model = "nomic-embed-text-v2-moe"
const ollamaEmbedURL = "http://localhost:11434/api/embeddings"

var client = &http.Client{Timeout: 60 * time.Second}

// Embed returns the embedding for text as a JSON-array string (e.g.
// "[0.1,0.2,...]"), suitable for storing in a sqlite column.
func Embed(ctx context.Context, text string) (string, error) {
	body, err := json.Marshal(map[string]string{"model": Model, "prompt": text})
	if err != nil {
		return "", fmt.Errorf("marshal embedding request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ollamaEmbedURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama not reachable: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return "", fmt.Errorf("ollama embedding request: %s", resp.Status)
	}
	var er struct {
		Embedding []float32 `json:"embedding"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&er); err != nil {
		return "", err
	}
	parts := make([]string, len(er.Embedding))
	for i, f := range er.Embedding {
		parts[i] = fmt.Sprintf("%g", f)
	}
	return "[" + strings.Join(parts, ",") + "]", nil
}
