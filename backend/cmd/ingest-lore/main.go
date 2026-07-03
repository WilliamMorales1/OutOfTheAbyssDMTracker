// ingest-lore downloads the "Out of the Abyss" adventure text from the
// 5etools data mirror, chunks it by section, embeds each chunk, and loads
// it into the chapter_chunks table backing Lore Search.
//
// Run manually (requires network access and a running Ollama instance):
//
//	go run ./cmd/ingest-lore
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"oota/internal/db"
	"oota/internal/embeddings"
)

const defaultSourceURL = "https://raw.githubusercontent.com/5etools-mirror-3/5etools-src/main/data/adventure/adventure-oota.json"
const maxChunkChars = 1000

type adventureFile struct {
	Data []section `json:"data"`
}

type section struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Entries []any  `json:"entries"`
}

var reTag = regexp.MustCompile(`\{@\w+ ([^}|]+)(\|[^}]*)?\}`)
var reWS = regexp.MustCompile(`\s+`)

func cleanText(s string) string {
	s = reTag.ReplaceAllString(s, "$1")
	s = reWS.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// flatten walks an arbitrary entries tree and returns plain-text paragraphs,
// each prefixed by the nearest heading name when one exists.
func flatten(entries []any, heading string, out *[]string) {
	for _, e := range entries {
		switch v := e.(type) {
		case string:
			if t := cleanText(v); t != "" {
				if heading != "" {
					*out = append(*out, heading+": "+t)
				} else {
					*out = append(*out, t)
				}
			}
		case map[string]any:
			typ, _ := v["type"].(string)
			if typ == "image" {
				continue
			}
			name, _ := v["name"].(string)
			h := heading
			if name != "" {
				h = name
			}
			if items, ok := v["items"].([]any); ok {
				flatten(items, h, out)
			}
			if rows, ok := v["rows"].([]any); ok {
				for _, row := range rows {
					if cells, ok := row.([]any); ok {
						var parts []string
						for _, c := range cells {
							if s, ok := c.(string); ok {
								parts = append(parts, cleanText(s))
							}
						}
						if len(parts) > 0 {
							*out = append(*out, h+" (table): "+strings.Join(parts, " | "))
						}
					}
				}
			}
			if nested, ok := v["entries"].([]any); ok {
				flatten(nested, h, out)
			}
		}
	}
}

// chunk groups consecutive paragraphs into ~maxChunkChars blocks.
func chunk(paragraphs []string) []string {
	var chunks []string
	var cur strings.Builder
	for _, p := range paragraphs {
		if cur.Len() > 0 && cur.Len()+len(p) > maxChunkChars {
			chunks = append(chunks, cur.String())
			cur.Reset()
		}
		if cur.Len() > 0 {
			cur.WriteString("\n\n")
		}
		cur.WriteString(p)
	}
	if cur.Len() > 0 {
		chunks = append(chunks, cur.String())
	}
	return chunks
}

func main() {
	sourceURL := flag.String("source", defaultSourceURL, "URL of the adventure JSON to ingest")
	dbPath := flag.String("db", "oota.db", "path to the sqlite database")
	flag.Parse()

	log.Printf("fetching %s", *sourceURL)
	resp, err := http.Get(*sourceURL)
	if err != nil {
		log.Fatalf("fetch: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("fetch: status %s", resp.Status)
	}
	var af adventureFile
	if err := json.NewDecoder(resp.Body).Decode(&af); err != nil {
		log.Fatalf("decode: %v", err)
	}

	conn, err := db.Open(*dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Exec(`DELETE FROM chapter_chunks`); err != nil {
		log.Fatalf("clear chapter_chunks: %v", err)
	}

	type job struct {
		chapterTitle string
		content      string
	}
	var jobs []job
	for _, sec := range af.Data {
		var paragraphs []string
		flatten(sec.Entries, "", &paragraphs)
		for _, c := range chunk(paragraphs) {
			jobs = append(jobs, job{sec.Name, c})
		}
	}
	log.Printf("embedding %d chunks", len(jobs))

	type result struct {
		job job
		emb string
		err error
	}
	resultCh := make(chan result)
	var wg sync.WaitGroup
	// Cap in-flight requests to GOMAXPROCS - the local Ollama server refuses
	// connections (not just slows down) if all chunks fire at once.
	sem := make(chan struct{}, runtime.GOMAXPROCS(0))
	for _, j := range jobs {
		wg.Add(1)
		go func(j job) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			emb, err := embeddings.Embed(context.Background(), j.content)
			resultCh <- result{j, emb, err}
		}(j)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	total := 0
	for res := range resultCh {
		if res.err != nil {
			log.Fatalf("embed chunk in %q: %v", res.job.chapterTitle, res.err)
		}
		if _, err := conn.Exec(
			`INSERT INTO chapter_chunks (chapter_title, content, embedding) VALUES (?, ?, ?)`,
			res.job.chapterTitle, res.job.content, res.emb,
		); err != nil {
			log.Fatalf("insert chunk in %q: %v", res.job.chapterTitle, err)
		}
		total++
		if total%25 == 0 {
			log.Printf("progress: %d/%d", total, len(jobs))
		}
	}
	log.Printf("done: %d chunks total", total)
}
