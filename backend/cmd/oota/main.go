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
	"time"

	"oota/internal/db"

	_ "golang.org/x/image/webp"
)

var conn *sql.DB
var roConn *sql.DB
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
	if err := rows.Err(); err != nil {
		return err
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
		if err := mrows.Err(); err != nil {
			mrows.Close()
			return err
		}
		if err := mrows.Close(); err != nil {
			return err
		}
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

const frontendDist = "../frontend"

func serveFrontend(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(frontendDist, filepath.Clean(r.URL.Path))
	if info, err := os.Stat(path); err == nil && !info.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.ServeFile(w, r, filepath.Join(frontendDist, "index.html"))
}

func main() {
	if err := db.RunMigrations("migrations", dbPath); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	var err error
	conn, err = db.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	roConn, err = db.OpenReadOnly(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer roConn.Close()
	q = db.New(conn)
	if err := loadGameMaps(ctx); err != nil {
		log.Fatalf("load maps: %v", err)
	}
	if err := syncNotesFromDisk(ctx); err != nil {
		log.Fatalf("sync notes: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	mux.HandleFunc("GET /api/monsters", handleAPIMonsters)
	mux.HandleFunc("GET /api/monsters/{id}", handleAPIMonster)
	mux.HandleFunc("GET /api/monster-stats", handleAPIMonsterStats)
	mux.HandleFunc("GET /api/spells", handleAPISpells)
	mux.HandleFunc("GET /api/spells/{id}", handleAPISpell)
	mux.HandleFunc("GET /api/sessions", handleAPISessions)
	mux.HandleFunc("GET /api/demon-lords", handleAPIDemonLords)
	mux.HandleFunc("GET /api/actions", handleAPIActions)
	mux.HandleFunc("GET /api/skill-areas", handleAPISkillAreas)
	mux.HandleFunc("GET /api/conditions", handleAPIConditions)
	mux.HandleFunc("GET /api/exhaustion-levels", handleAPIExhaustionLevels)
	mux.HandleFunc("GET /api/maps", handleAPIMaps)
	mux.HandleFunc("GET /api/chat", handleAPIChat)
	mux.HandleFunc("GET /api/ollama-models", handleAPIOllamaModels)
	mux.HandleFunc("GET /api/search", handleAPISearch)
	mux.HandleFunc("GET /api/notes", handleAPINotesList)
	mux.HandleFunc("GET /api/notes/{name}", handleAPIGetNote)
	mux.HandleFunc("PUT /api/notes/{name}", handleAPIPutNote)
	mux.HandleFunc("GET /api/initiative-presets", handleAPIInitiativePresets)
	mux.HandleFunc("PUT /api/initiative-presets/{name}", handleAPIPutInitiativePreset)
	mux.HandleFunc("DELETE /api/initiative-presets/{name}", handleAPIDeleteInitiativePreset)
	mux.HandleFunc("/", serveFrontend)

	server := &http.Server{
		Addr:              "localhost:8080",
		Handler:           logRequests(mux),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       2 * time.Minute,
	}
	log.Println("Listening on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
