package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

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
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Cr       string `json:"cr"`
	ImageUrl string `json:"imageUrl"`
}

func monsterToDTO(m db.ListMonstersRow) monsterDTO {
	return monsterDTO{
		ID:       m.ID,
		Name:     m.Name,
		Type:     m.Type.String,
		Cr:       m.Cr.String,
		ImageUrl: m.ImageUrl,
	}
}

func handleAPIMonsters(w http.ResponseWriter, r *http.Request) {
	listHandler(q.ListMonsters, monsterToDTO)(w, r)
}

// monsterStatDTO is the lightweight ac/hp/dex lookup used to autofill the
// initiative tracker, kept separate from the full bestiary list/detail DTOs.
type monsterStatDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Ac   int64  `json:"ac"`
	Hp   int64  `json:"hp"`
	Dex  int64  `json:"dex"`
}

func monsterStatToDTO(m db.ListMonsterStatsRow) monsterStatDTO {
	return monsterStatDTO{
		ID:   m.ID,
		Name: m.Name,
		Ac:   m.Ac.Int64,
		Hp:   m.Hp.Int64,
		Dex:  m.Dex.Int64,
	}
}

func handleAPIMonsterStats(w http.ResponseWriter, r *http.Request) {
	listHandler(q.ListMonsterStats, monsterStatToDTO)(w, r)
}

// statBlockEntry is one named trait/action/reaction/legendary-action block.
type statBlockEntry struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func parseStatBlockEntries(jsonText string) []statBlockEntry {
	if jsonText == "" {
		return nil
	}
	var entries []statBlockEntry
	if err := json.Unmarshal([]byte(jsonText), &entries); err != nil {
		return nil
	}
	return entries
}

type monsterDetailDTO struct {
	ID                  int64            `json:"id"`
	Name                string           `json:"name"`
	Type                string           `json:"type"`
	Size                string           `json:"size"`
	Alignment           string           `json:"alignment"`
	Cr                  string           `json:"cr"`
	Source              string           `json:"source"`
	Hp                  int64            `json:"hp"`
	HpFormula           string           `json:"hpFormula"`
	Ac                  int64            `json:"ac"`
	AcDesc              string           `json:"acDesc"`
	Speed               string           `json:"speed"`
	Str                 int64            `json:"str"`
	Dex                 int64            `json:"dex"`
	Con                 int64            `json:"con"`
	Int                 int64            `json:"int"`
	Wis                 int64            `json:"wis"`
	Cha                 int64            `json:"cha"`
	SavingThrows        string           `json:"savingThrows"`
	Skills              string           `json:"skills"`
	DamageResistances   string           `json:"damageResistances"`
	DamageImmunities    string           `json:"damageImmunities"`
	Vulnerabilities     string           `json:"vulnerabilities"`
	ConditionImmunities string           `json:"conditionImmunities"`
	Senses              string           `json:"senses"`
	PassivePerception   int64            `json:"passivePerception"`
	Languages           string           `json:"languages"`
	Environment         string           `json:"environment"`
	ImageUrl            string           `json:"imageUrl"`
	Traits              []statBlockEntry `json:"traits"`
	Actions             []statBlockEntry `json:"actions"`
	Reactions           []statBlockEntry `json:"reactions"`
	LegendaryActions    []statBlockEntry `json:"legendaryActions"`
	Spellcasting        []statBlockEntry `json:"spellcasting"`
	Notes               string           `json:"notes"`
}

func monsterDetailToDTO(m db.GetMonsterRow) monsterDetailDTO {
	return monsterDetailDTO{
		ID:                  m.ID,
		Name:                m.Name,
		Type:                m.Type.String,
		Size:                m.Size,
		Alignment:           m.Alignment,
		Cr:                  m.Cr.String,
		Source:              m.Source,
		Hp:                  m.Hp.Int64,
		HpFormula:           m.HpFormula.String,
		Ac:                  m.Ac.Int64,
		AcDesc:              m.AcDesc.String,
		Speed:               m.Speed.String,
		Str:                 m.Str.Int64,
		Dex:                 m.Dex.Int64,
		Con:                 m.Con.Int64,
		Int:                 m.IntScore.Int64,
		Wis:                 m.Wis.Int64,
		Cha:                 m.Cha.Int64,
		SavingThrows:        m.SavingThrows,
		Skills:              m.Skills,
		DamageResistances:   m.DamageResistances,
		DamageImmunities:    m.DamageImmunities,
		Vulnerabilities:     m.Vulnerabilities,
		ConditionImmunities: m.ConditionImmunities,
		Senses:              m.Senses,
		PassivePerception:   m.PassivePerception.Int64,
		Languages:           m.Languages,
		Environment:         m.Environment,
		ImageUrl:            m.ImageUrl,
		Traits:              parseStatBlockEntries(m.Traits),
		Actions:             parseStatBlockEntries(m.Actions),
		Reactions:           parseStatBlockEntries(m.Reactions),
		LegendaryActions:    parseStatBlockEntries(m.LegendaryActions),
		Spellcasting:        parseStatBlockEntries(m.Spellcasting),
		Notes:               m.Notes,
	}
}

func handleAPIMonster(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/monsters/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid monster id", 400)
		return
	}
	m, err := q.GetMonster(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	writeJSON(w, monsterDetailToDTO(m))
}

func handleAPIMaps(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, gameMaps)
}

func handleAPIRefs(w http.ResponseWriter, r *http.Request) {
	rows, err := conn.QueryContext(r.Context(), `SELECT id, title, content FROM Refs ORDER BY id`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()
	type refDTO struct {
		ID      string          `json:"id"`
		Title   string          `json:"title"`
		Content json.RawMessage `json:"content"`
	}
	out := []refDTO{}
	for rows.Next() {
		var ref refDTO
		var content string
		if err := rows.Scan(&ref.ID, &ref.Title, &content); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		ref.Content = json.RawMessage(content)
		out = append(out, ref)
	}
	writeJSON(w, out)
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
	ctx, cancel := context.WithTimeout(r.Context(), 90*time.Second)
	defer cancel()
	answer, err := runAgent(ctx, q)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	writeJSON(w, struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}{q, answer})
}

var reFTSTerm = regexp.MustCompile(`[A-Za-z0-9']+`)

// ftsMatchQuery turns free text into an FTS5 MATCH expression. Each term is
// prefix-matched and OR'd together so partial/misspelled words still surface
// results, the same tolerant-matching behavior search engines use.
func ftsMatchQuery(query string) string {
	terms := reFTSTerm.FindAllString(query, -1)
	if len(terms) == 0 {
		return ""
	}
	quoted := make([]string, len(terms))
	for i, t := range terms {
		quoted[i] = `"` + strings.ReplaceAll(t, `"`, `""`) + `"*`
	}
	return strings.Join(quoted, " OR ")
}

// rrfConst is the standard reciprocal-rank-fusion constant used by hybrid
// search systems (e.g. Elasticsearch, Azure AI Search) to blend independently
// ranked result lists - here, BM25 keyword relevance and embedding similarity.
const rrfConst = 60

func searchLore(ctx context.Context, query string) ([]searchResult, error) {
	if query == "" {
		return []searchResult{}, nil
	}

	type chunk struct {
		chapterTitle string
		content      string
	}
	fused := map[int]float64{}
	chunks := map[int]chunk{}

	// Keyword relevance: SQLite FTS5 BM25 ranking.
	if matchQ := ftsMatchQuery(query); matchQ != "" {
		rows, err := conn.QueryContext(ctx,
			`SELECT rowid, chapter_title, content FROM chapter_chunks_fts
			 WHERE chapter_chunks_fts MATCH ? ORDER BY bm25(chapter_chunks_fts) LIMIT 25`, matchQ)
		if err != nil {
			return nil, fmt.Errorf("search error: %w", err)
		}
		rank := 0
		for rows.Next() {
			var id int
			var c chunk
			if err := rows.Scan(&id, &c.chapterTitle, &c.content); err != nil {
				rows.Close()
				return nil, err
			}
			rank++
			fused[id] += 1.0 / float64(rrfConst+rank)
			chunks[id] = c
		}
		rows.Close()
	}

	// Semantic relevance: embedding cosine similarity.
	if emb, err := queryEmbedding(ctx, query); err == nil {
		rows, err := conn.QueryContext(ctx, `SELECT id, chapter_title, content, embedding FROM chapter_chunks`)
		if err != nil {
			return nil, fmt.Errorf("search error: %w", err)
		}
		type scored struct {
			id    int
			chunk chunk
			score float64
		}
		var semantic []scored
		for rows.Next() {
			var id int
			var c chunk
			var embeddingJSON string
			if err := rows.Scan(&id, &c.chapterTitle, &c.content, &embeddingJSON); err != nil {
				rows.Close()
				return nil, err
			}
			score, err := cosineSimilarity(emb, embeddingJSON)
			if err != nil {
				continue
			}
			semantic = append(semantic, scored{id, c, score})
		}
		rows.Close()
		sort.Slice(semantic, func(i, j int) bool { return semantic[i].score > semantic[j].score })
		if len(semantic) > 25 {
			semantic = semantic[:25]
		}
		for rank, s := range semantic {
			fused[s.id] += 1.0 / float64(rrfConst+rank+1)
			chunks[s.id] = s.chunk
		}
	}

	type idScore struct {
		id    int
		score float64
	}
	ranked := make([]idScore, 0, len(fused))
	for id, score := range fused {
		ranked = append(ranked, idScore{id, score})
	}
	sort.Slice(ranked, func(i, j int) bool { return ranked[i].score > ranked[j].score })
	if len(ranked) > 5 {
		ranked = ranked[:5]
	}

	results := []searchResult{}
	maxScore := 0.0
	if len(ranked) > 0 {
		maxScore = ranked[0].score
	}
	for _, rs := range ranked {
		c := chunks[rs.id]
		norm := 0.0
		if maxScore > 0 {
			norm = rs.score / maxScore
		}
		results = append(results, searchResult{ChapterTitle: c.chapterTitle, Content: c.content, Score: norm})
	}

	return results, nil
}

func handleAPISearch(w http.ResponseWriter, r *http.Request) {
	results, err := searchLore(r.Context(), r.URL.Query().Get("q"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	writeJSON(w, results)
}

var validNoteName = regexp.MustCompile(`^[A-Za-z0-9_-]+\.md$`)

const notesDir = "notes"

func syncNotesFromDisk(ctx context.Context) error {
	files, err := filepath.Glob(filepath.Join(notesDir, "*.md"))
	if err != nil {
		return err
	}
	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		name := filepath.Base(f)
		if err := q.UpsertNote(ctx, db.UpsertNoteParams{Name: name, Content: string(content)}); err != nil {
			return err
		}
	}
	return nil
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
		if err := os.WriteFile(filepath.Join(notesDir, name), body, 0644); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(204)

	default:
		http.Error(w, "method not allowed", 405)
	}
}
