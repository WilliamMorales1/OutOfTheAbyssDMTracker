package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"oota/internal/db"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// listHandler builds a GET handler that fetches rows and maps each to a DTO.
func listHandler[T, R any](list func(ctx context.Context) ([]T, error), toDTO func(T) R) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := list(r.Context())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		out := make([]R, len(rows))
		for i, row := range rows {
			out[i] = toDTO(row)
		}
		writeJSON(w, out)
	}
}

func stars(n int64) string {
	if n < 0 {
		n = 0
	}
	if n > 5 {
		n = 5
	}
	return strings.Repeat("★", int(n)) + strings.Repeat("☆", 5-int(n))
}

type sessionDTO struct {
	SessionNum    int64  `json:"sessionNum"`
	Title         string `json:"title"`
	Chapters      string `json:"chapters"`
	LevelStart    int64  `json:"levelStart"`
	LevelEnd      int64  `json:"levelEnd"`
	Summary       string `json:"summary"`
	KeyEncounters string `json:"keyEncounters"`
	KeyNpcs       string `json:"keyNpcs"`
	Checkpoint    string `json:"checkpoint"`
}

func sessionToDTO(s db.Session) sessionDTO {
	return sessionDTO{
		SessionNum:    s.SessionNum,
		Title:         s.Title,
		Chapters:      s.Chapters.String,
		LevelStart:    s.LevelStart.Int64,
		LevelEnd:      s.LevelEnd.Int64,
		Summary:       s.Summary.String,
		KeyEncounters: s.KeyEncounters.String,
		KeyNpcs:       s.KeyNpcs.String,
		Checkpoint:    s.Checkpoint.String,
	}
}

func handleAPISessions(w http.ResponseWriter, r *http.Request) {
	listHandler(q.ListSessions, sessionToDTO)(w, r)
}

type npcDTO struct {
	Name         string `json:"name"`
	Madness      int64  `json:"madness"`
	MadnessStars string `json:"madnessStars"`
	Disposition  string `json:"disposition"`
	Location     string `json:"location"`
	Description  string `json:"description"`
	Notes        string `json:"notes"`
}

func npcToDTO(n db.ListNPCsRow) npcDTO {
	return npcDTO{
		Name:         n.Name.String,
		Madness:      n.Madness.Int64,
		MadnessStars: stars(n.Madness.Int64),
		Disposition:  n.Disposition.String,
		Location:     n.Location.String,
		Description:  n.Description.String,
		Notes:        n.Notes.String,
	}
}

func handleAPINPCs(w http.ResponseWriter, r *http.Request) {
	listHandler(q.ListNPCs, npcToDTO)(w, r)
}

type monsterDTO struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Cr        string `json:"cr"`
	Hp        int64  `json:"hp"`
	HpFormula string `json:"hpFormula"`
	Ac        int64  `json:"ac"`
	AcDesc    string `json:"acDesc"`
	Speed     string `json:"speed"`
	Dex       int64  `json:"dex"`
}

func monsterToDTO(m db.ListMonstersRow) monsterDTO {
	return monsterDTO{
		Name:      m.Name,
		Type:      m.Type.String,
		Cr:        m.Cr.String,
		Hp:        m.Hp.Int64,
		HpFormula: m.HpFormula.String,
		Ac:        m.Ac.Int64,
		AcDesc:    m.AcDesc.String,
		Speed:     m.Speed.String,
		Dex:       m.Dex.Int64,
	}
}

func handleAPIMonsters(w http.ResponseWriter, r *http.Request) {
	listHandler(q.ListMonsters, monsterToDTO)(w, r)
}

func handleAPIMaps(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, gameMaps)
}

func handleAPIChat(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		q = r.FormValue("q")
	}
	if q == "" {
		http.Error(w, "missing q", 400)
		return
	}
	answer, err := runAgent(r.Context(), q)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	writeJSON(w, struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}{q, answer})
}

func handleAPISearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		writeJSON(w, []searchResult{})
		return
	}

	emb, err := queryEmbedding(r.Context(), query)
	if err != nil {
		http.Error(w, "Embedding error: "+err.Error(), 500)
		return
	}

	rows, err := conn.QueryContext(r.Context(), `SELECT chapter_title, content, embedding FROM chapter_chunks`)
	if err != nil {
		http.Error(w, "Search error: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	var results []searchResult
	for rows.Next() {
		var chapterTitle, content, embeddingJSON string
		if err := rows.Scan(&chapterTitle, &content, &embeddingJSON); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		score, err := cosineSimilarity(emb, embeddingJSON)
		if err != nil {
			continue
		}
		results = append(results, searchResult{ChapterTitle: chapterTitle, Content: content, Score: score})
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Score > results[j].Score })
	if len(results) > 5 {
		results = results[:5]
	}
	if results == nil {
		results = []searchResult{}
	}

	writeJSON(w, results)
}

var validNoteName = regexp.MustCompile(`^[A-Za-z0-9_-]+\.md$`)

const notesSeedFile = "migrations/002_seed_data.up.sql"
const notesSeedMarker = "-- Notes\n"

func sqlQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}

// regenerateNotesSeed rewrites the Notes block in the seed migration so a
// reseed restores the latest saved note content.
func regenerateNotesSeed(ctx context.Context) error {
	notes, err := q.ListNotes(ctx)
	if err != nil {
		return err
	}

	existing, err := os.ReadFile(notesSeedFile)
	if err != nil {
		return err
	}
	idx := strings.Index(string(existing), notesSeedMarker)
	if idx == -1 {
		return fmt.Errorf("%s: missing %q marker", notesSeedFile, notesSeedMarker)
	}

	var b strings.Builder
	b.Write(existing[:idx])
	b.WriteString(notesSeedMarker)
	for _, n := range notes {
		fmt.Fprintf(&b, "INSERT INTO Notes (name, content) VALUES (%s, %s);\n", sqlQuote(n.Name), sqlQuote(n.Content))
	}

	return os.WriteFile(notesSeedFile, []byte(b.String()), 0644)
}

func handleAPINotesList(w http.ResponseWriter, r *http.Request) {
	names, err := q.ListNoteNames(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	writeJSON(w, names)
}

func handleAPINote(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/notes/")
	if !validNoteName.MatchString(name) {
		http.Error(w, "invalid note name", 400)
		return
	}

	switch r.Method {
	case http.MethodGet:
		note, err := q.GetNote(r.Context(), name)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		writeJSON(w, struct {
			Name    string `json:"name"`
			Content string `json:"content"`
		}{note.Name, note.Content})

	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := q.UpsertNote(r.Context(), db.UpsertNoteParams{Name: name, Content: string(body)}); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := regenerateNotesSeed(r.Context()); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(204)

	default:
		http.Error(w, "method not allowed", 405)
	}
}
