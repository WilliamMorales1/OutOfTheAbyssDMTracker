package main

import (
	"context"
	"database/sql"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"oota/internal/db"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "golang.org/x/image/webp"
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
	Markers []Marker `json:"markers"`
}

func imgVB(imgPath string) string {
	f, err := os.Open(imgPath)
	if err != nil {
		return ""
	}
	defer f.Close()
	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("0 0 %d %d", cfg.Width, cfg.Height)
}

func loadGameMaps(ctx context.Context) error {
	rows, err := conn.QueryContext(ctx, `SELECT id, img FROM GameMaps ORDER BY id`)
	if err != nil {
		return err
	}
	defer rows.Close()
	maps := []GameMap{}
	for rows.Next() {
		gm := GameMap{Markers: []Marker{}}
		if err := rows.Scan(&gm.ID, &gm.Img); err != nil {
			return err
		}
		gm.VB = imgVB(strings.TrimPrefix(gm.Img, "./"))
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
	if err := syncNotesFromDisk(ctx); err != nil {
		log.Fatalf("sync notes: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	mux.HandleFunc("/api/npcs", handleAPINPCs)
	mux.HandleFunc("/api/monsters", handleAPIMonsters)
	mux.HandleFunc("/api/monsters/", handleAPIMonster)
	mux.HandleFunc("/api/monster-stats", handleAPIMonsterStats)
	mux.HandleFunc("/api/sessions", handleAPISessions)
	mux.HandleFunc("/api/maps", handleAPIMaps)
	mux.HandleFunc("/api/refs", handleAPIRefs)
	mux.HandleFunc("/api/chat", handleAPIChat)
	mux.HandleFunc("/api/search", handleAPISearch)
	mux.HandleFunc("/api/notes", handleAPINotesList)
mux.HandleFunc("/api/notes/", handleAPINote)
	mux.HandleFunc("/", serveFrontend)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", logRequests(mux)))
}
