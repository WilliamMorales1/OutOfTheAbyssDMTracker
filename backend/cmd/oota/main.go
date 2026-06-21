package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"oota/internal/db"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

var conn *sql.DB
var q *db.Queries
var gameMaps []GameMap

type Marker struct {
	I     int    `json:"i"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GameMap struct {
	ID      string   `json:"id"`
	Img     string   `json:"img"`
	VB      string   `json:"vb"`
	W       string   `json:"w"`
	H       string   `json:"h"`
	Markers []Marker `json:"markers"`
}

func vbDims(vb string) (w, h string) {
	parts := strings.Fields(vb)
	if len(parts) == 4 {
		return parts[2], parts[3]
	}
	return "100%", "100%"
}

func loadGameMaps(ctx context.Context) error {
	rows, err := conn.QueryContext(ctx, `SELECT id, img, vb FROM GameMaps ORDER BY id`)
	if err != nil {
		return err
	}
	defer rows.Close()
	maps := []GameMap{}
	for rows.Next() {
		gm := GameMap{Markers: []Marker{}}
		if err := rows.Scan(&gm.ID, &gm.Img, &gm.VB); err != nil {
			return err
		}
		maps = append(maps, gm)
	}
	for i, gm := range maps {
		mrows, err := conn.QueryContext(ctx, `SELECT i, x, y, title, body FROM MapMarkers WHERE map_id=? ORDER BY i`, gm.ID)
		if err != nil {
			return err
		}
		for mrows.Next() {
			var m Marker
			if err := mrows.Scan(&m.I, &m.X, &m.Y, &m.Title, &m.Body); err != nil {
				mrows.Close()
				return err
			}
			maps[i].Markers = append(maps[i].Markers, m)
		}
		mrows.Close()
	}
	for i, gm := range maps {
		maps[i].W, maps[i].H = vbDims(gm.VB)
	}
	gameMaps = maps
	return nil
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

const dbPath = "oota.db"
const dbURL = "sqlite://" + dbPath + "?_pragma=foreign_keys(1)"

func runMigrations() {
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatalf("migrations init: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrations up: %v", err)
	}
}

const frontendDist = "../frontend/dist"

func serveFrontend(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(frontendDist, filepath.Clean(r.URL.Path))
	if info, err := os.Stat(path); err == nil && !info.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.ServeFile(w, r, filepath.Join(frontendDist, "index.html"))
}

func main() {
	runMigrations()

	ctx := context.Background()
	var err error
	conn, err = sql.Open("sqlite", dbPath+"?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	q = db.New(conn)
	if err := loadGameMaps(ctx); err != nil {
		log.Fatalf("load maps: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	mux.HandleFunc("/api/npcs", handleAPINPCs)
	mux.HandleFunc("/api/monsters", handleAPIMonsters)
	mux.HandleFunc("/api/monsters/", handleAPIMonster)
	mux.HandleFunc("/api/monster-stats", handleAPIMonsterStats)
	mux.HandleFunc("/api/sessions", handleAPISessions)
	mux.HandleFunc("/api/maps", handleAPIMaps)
	mux.HandleFunc("/api/chat", handleAPIChat)
	mux.HandleFunc("/api/search", handleAPISearch)
	mux.HandleFunc("/api/notes", handleAPINotesList)
	mux.HandleFunc("/api/notes/", handleAPINote)
	mux.HandleFunc("/", serveFrontend)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", logRequests(mux)))
}
