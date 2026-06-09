package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"oota/db"

	"github.com/a-h/templ"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn
var q *db.Queries

type encounterRow struct {
	db.ListEncountersRow
	Monsters []db.ListEncounterMonstersRow
}

func renderWith[T any](w http.ResponseWriter, r *http.Request, query func(context.Context) (T, error), tmpl func(T) templ.Component) {
	data, err := query(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl(data).Render(r.Context(), w)
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

const dbURL = "postgres://wsm52:H&pg@localhost/oota?sslmode=disable"

func runMigrations() {
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatalf("migrations init: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrations up: %v", err)
	}
}

func main() {
	runMigrations()

	ctx := context.Background()
	var err error
	conn, err = pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	q = db.New(conn)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "oota.html") })
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	mux.HandleFunc("/panel/", handlePanel)
	mux.HandleFunc("/locations", handleLocations)
	mux.HandleFunc("/npcs", handleNPCs)
	mux.HandleFunc("/encounters", handleEncounters)
	mux.HandleFunc("/events", handleEvents)
	mux.HandleFunc("/monsters", handleMonsters)
	mux.HandleFunc("/sessions", handleSessions)
	mux.HandleFunc("/maps", handleMaps)
	mux.HandleFunc("/chat", handleChat)
	mux.HandleFunc("/search", handleSearch)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", logRequests(mux)))
}

func handlePanel(w http.ResponseWriter, r *http.Request) {
	tab := strings.TrimPrefix(r.URL.Path, "/panel/")
	ctx := r.Context()

	w.Header().Set("Content-Type", "text/html")
	PanelTabs(tab).Render(ctx, w)

	switch tab {
	case "chat":
		ChatPanel().Render(ctx, w)
	case "search":
		SearchPanel().Render(ctx, w)
	case "maps":
		MapsPanel().Render(ctx, w)
	default:
		DefaultPanel("/"+tab).Render(ctx, w)
	}
}

func handleMaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	MapsPanel().Render(r.Context(), w)
}

func handleSessions(w http.ResponseWriter, r *http.Request) {
	renderWith(w, r, q.ListSessions, SessionsTable)
}

func handleLocations(w http.ResponseWriter, r *http.Request) {
	renderWith(w, r, q.ListLocations, LocationsTable)
}

func handleNPCs(w http.ResponseWriter, r *http.Request) {
	renderWith(w, r, q.ListNPCs, NPCsTable)
}

func handleEncounters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rows, err := q.ListEncounters(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	encs := make([]encounterRow, len(rows))
	idIndex := map[int64]int{}
	for i, e := range rows {
		encs[i] = encounterRow{ListEncountersRow: e}
		idIndex[e.ID] = i
	}

	mRows, err := q.ListEncounterMonsters(ctx)
	if err != nil {
		log.Printf("encounter monsters query: %v", err)
	} else {
		for _, m := range mRows {
			if idx, ok := idIndex[m.EncounterID]; ok {
				encs[idx].Monsters = append(encs[idx].Monsters, m)
			}
		}
	}

	w.Header().Set("Content-Type", "text/html")
	EncountersTable(encs).Render(ctx, w)
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	renderWith(w, r, q.ListEvents, EventsTable)
}

func handleMonsters(w http.ResponseWriter, r *http.Request) {
	renderWith(w, r, q.ListMonsters, MonstersTable)
}

func dangerStars(d int32) string {
	if d < 0 {
		d = 0
	}
	filled := min(d, 5)
	return strings.Repeat("★", int(filled)) + strings.Repeat("☆", int(5-filled))
}
