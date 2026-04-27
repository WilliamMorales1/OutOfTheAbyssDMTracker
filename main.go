package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	ctx := context.Background()
	var err error
	db, err = pgx.Connect(ctx, "postgres://wsm52:H&pg@localhost/oota")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	initAI()

	if err := openIndex(); err != nil {
		log.Fatal("bleve open:", err)
	}

	go func() {
		if err := indexAll(); err != nil {
			log.Println("Bleve index error:", err)
		} else {
			log.Println("Bleve indexed.")
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveHTML)
	mux.HandleFunc("/panel/", handlePanel)
	mux.HandleFunc("/locations", handleLocations)
	mux.HandleFunc("/npcs", handleNPCs)
	mux.HandleFunc("/encounters", handleEncounters)
	mux.HandleFunc("/events", handleEvents)
	mux.HandleFunc("/chat", handleChat)
	mux.HandleFunc("/reindex", handleReindex)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", logRequests(mux)))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "oota.html")
}

func handlePanel(w http.ResponseWriter, r *http.Request) {
	tab := strings.TrimPrefix(r.URL.Path, "/panel/")

	tabs := []struct {
		Name string
		Path string
	}{
		{"Locations", "locations"},
		{"NPCs", "npcs"},
		{"Encounters", "encounters"},
		{"Events", "events"},
		{"Ask Agent", "chat"},
	}

	var tabsHTML strings.Builder
	for _, t := range tabs {
		cls := "tab"
		if t.Path == tab {
			cls = "tab active"
		}
		fmt.Fprintf(&tabsHTML, `<button class=%q hx-get="/panel/%s" hx-target="#panel" hx-swap="innerHTML">%s</button>`,
			cls, t.Path, t.Name)
	}
	oobTabs := fmt.Sprintf(`<div id="tabs" hx-swap-oob="true">%s</div>`, tabsHTML.String())

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, oobTabs)

	switch tab {
	case "chat":
		fmt.Fprint(w, `
<button hx-post="/reindex" hx-target="#reindex-status" hx-swap="innerHTML">Re-index DB →</button>
<span id="reindex-status"></span>
<div id="chat-history"></div>
<form hx-post="/chat" 
      hx-target="#chat-history" 
      hx-swap="beforeend" 
      hx-on::after-request="this.reset()" 
      hx-indicator="#chat-spinner"
      class="mt-3">
  
  <div class="input-group">
    <input name="q" 
           type="text" 
           class="form-control bg-dark text-light border-secondary" 
           placeholder="Ask anything about the campaign..." 
           required>
    <button class="btn btn-primary" type="submit">Ask</button>
  </div>

  <div class="mt-2">
    <span id="chat-spinner" class="htmx-indicator spinner-border spinner-border-sm text-warning" role="status"></span>
    <span id="chat-spinner-text" class="htmx-indicator ms-2 text-secondary small">Thinking...</span>
  </div>
</form>`)
	default:
		endpoint := "/" + tab
		fmt.Fprintf(w, `
<form hx-get="%s" hx-target="#results" hx-swap="innerHTML">
  <input name="q" type="text" placeholder="Search...">
  <button type="submit">Search</button>
  <span class="htmx-indicator" id="spinner">Loading...</span>
</form>
<div id="results" hx-get="%s" hx-trigger="load" hx-swap="innerHTML"></div>`, endpoint, endpoint)
	}
}

func handleLocations(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, type, status, danger, description, secrets FROM Locations ORDER BY name`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var locs []location
	for rows.Next() {
		var l location
		if err := rows.Scan(&l.Id, &l.Name, &l.Type_, &l.Status, &l.Danger, &l.Description, &l.Secrets); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		locs = append(locs, l)
	}

	renderTable(w, locationsTmpl, locs)
}

func handleNPCs(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, madness, disposition, location, notes, description FROM NPCS ORDER BY name`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var npcs []npc
	for rows.Next() {
		var n npc
		if err := rows.Scan(&n.Id, &n.Name, &n.Madness, &n.Disposition, &n.Location, &n.Notes, &n.Description); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		npcs = append(npcs, n)
	}

	renderTable(w, npcsTmpl, npcs)
}

func handleEncounters(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, COALESCE(location,''), difficulty, status, enemies, levelup, notes FROM Encounters ORDER BY name`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var encs []encounter
	for rows.Next() {
		var e encounter
		if err := rows.Scan(&e.Id, &e.Name, &e.Location, &e.Difficulty, &e.Status, &e.Enemies, &e.Levelup, &e.Notes); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		encs = append(encs, e)
	}

	renderTable(w, encountersTmpl, encs)
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, title, category, description FROM Events ORDER BY title`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var evts []event
	for rows.Next() {
		var e event
		if err := rows.Scan(&e.Id, &e.Title, &e.Category, &e.Description); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		evts = append(evts, e)
	}

	renderTable(w, eventsTmpl, evts)
}

func renderTable(w http.ResponseWriter, tmplStr string, data interface{}) {
	tmpl, err := template.New("t").Funcs(template.FuncMap{
		"danger": func(d int) string {
			if d < 0 {
				d = 0
			}
			filled := d
			if filled > 5 {
				filled = 5
			}
			return strings.Repeat("★", filled) + strings.Repeat("☆", 5-filled)
		},
	}).Parse(tmplStr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		fmt.Fprintf(w, "<p class='error'>%s</p>", err.Error())
	}
}

var locationsTmpl = `
<div class="table-responsive">
  <table class="table table-dark table-hover table-bordered datatable w-100">
    <thead>
      <tr>
        <th>Name</th>
        <th>Type</th>
        <th>Status</th>
        <th>Danger</th>
        <th>Description</th>
        <th>Secrets</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
      <tr>
        <td><strong>{{.Name}}</strong></td>
        <td>{{.Type_}}</td>
        <td><span class="badge bg-primary">{{.Status}}</span></td>
        <td data-order="{{.Danger}}">{{danger .Danger}}</td>
        <td>{{.Description}}</td>
        <td>{{.Secrets}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
</div>`

var npcsTmpl = `
<div class="table-responsive">
  {{if .}}
  <table class="table table-dark table-hover table-bordered datatable w-100">
    <thead>
      <tr>
        <th>Name</th>
        <th>Madness</th>
        <th>Disposition</th>
        <th>Location</th>
        <th>Description</th>
        <th>Notes</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
      <tr>
        <td><strong>{{.Name}}</strong></td>
        <td class="danger" data-order="{{.Madness}}">{{danger .Madness}}</td>
        <td><span class="badge disp-{{.Disposition}}">{{.Disposition}}</span></td>
        <td>{{.Location}}</td>
        <td>{{.Description}}</td>
        <td>{{.Notes}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
  {{else}}<p class="empty">No NPCs found.</p>
  {{end}}
</div>`

var encountersTmpl = `
<div class="table-responsive">
  {{if .}}
  <table class="table table-dark table-hover table-bordered datatable w-100">
    <thead>
      <tr>
        <th>Name</th>
        <th>Location</th>
        <th>Difficulty</th>
        <th>Status</th>
        <th>Enemies</th>
        <th>Level Up</th>
        <th>Notes</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
      <tr>
        <td><strong>{{.Name}}</strong></td>
        <td>{{.Location}}</td>
        <td class="danger" data-order="{{.Difficulty}}">{{danger .Difficulty}}</td>
        <td><span class="badge status-{{.Status}}">{{.Status}}</span></td>
        <td>{{.Enemies}}</td>
        <td>{{.Levelup}}</td>
        <td>{{.Notes}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
  {{else}}
    <p class="empty p-3">No encounters found.</p>
  {{end}}
</div>`

var eventsTmpl = `
<div class="table-responsive">
  {{if .}}
  <table class="table table-dark table-hover table-bordered datatable w-100">
    <thead>
      <tr>
        <th>Title</th>
        <th>Category</th>
        <th>Description</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
      <tr>
        <td><strong>{{.Title}}</strong></td>
        <td><span class="badge cat-{{.Category}}">{{.Category}}</span></td>
        <td>{{.Description}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
  {{else}}<p class="empty">No events found.</p>
  {{end}}
</div>`

type location struct {
	Id          int64
	Name        string
	Type_       string
	Status      string
	Danger      int
	Description string
	Secrets     string
}

type npc struct {
	Id          int64
	Name        string
	Madness     int
	Disposition string
	Location    string
	Notes       string
	Description string
}

type encounter struct {
	Id         int64
	Name       string
	Location   string
	Difficulty int
	Status     string
	Enemies    string
	Levelup    bool
	Notes      string
}

type event struct {
	Id          int64
	Title       string
	Category    string
	Description string
}
