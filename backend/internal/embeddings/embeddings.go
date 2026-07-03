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
)

const Model = "nomic-embed-text-v2-moe"
const ollamaEmbedURL = "http://localhost:11434/api/embeddings"

// Embed returns the embedding for text as a JSON-array string (e.g.
// "[0.1,0.2,...]"), suitable for storing in a sqlite column.
func Embed(ctx context.Context, text string) (string, error) {
	body, _ := json.Marshal(map[string]string{"model": Model, "prompt": text})
	req, err := http.NewRequestWithContext(ctx, "POST", ollamaEmbedURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama not reachable: %w", err)
	}
	defer resp.Body.Close()
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
