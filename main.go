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

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveHTML)
	mux.HandleFunc("/panel/", handlePanel)
	mux.HandleFunc("/locations", handleLocations)
	mux.HandleFunc("/npcs", handleNPCs)
	mux.HandleFunc("/encounters", handleEncounters)
	mux.HandleFunc("/events", handleEvents)
	mux.HandleFunc("/monsters", handleMonsters)
	mux.HandleFunc("/sessions", handleSessions)
	mux.HandleFunc("/chat", handleChat)
	mux.HandleFunc("/search", handleSearch)

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
		{"Sessions", "sessions"},
		{"Locations", "locations"},
		{"NPCs", "npcs"},
		{"Encounters", "encounters"},
		{"Monsters", "monsters"},
		{"Events", "events"},
		{"Ask Agent", "chat"},
		{"Lore Search", "search"},
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
	case "search":
		fmt.Fprint(w, `
<form hx-get="/search" hx-target="#results" hx-swap="innerHTML" hx-indicator="#search-spinner">
  <div class="input-group mb-3">
    <input name="q" type="text" class="form-control bg-dark text-light border-secondary"
           placeholder="Search lore semantically... (e.g. 'drow priestess tactics')" required>
    <button class="btn btn-warning" type="submit">Search</button>
  </div>
  <span id="search-spinner" class="htmx-indicator spinner-border spinner-border-sm text-warning me-2"></span>
  <span class="htmx-indicator text-secondary small">Searching...</span>
</form>
<div id="results" class="mt-3"></div>`)
	default:
		endpoint := "/" + tab
		fmt.Fprintf(w, `
<div id="results" hx-get="%s" hx-trigger="load" hx-swap="innerHTML"></div>`, endpoint)
	}
}

func handleSessions(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, session_num, title, chapters, level_start, level_end, summary, key_encounters, key_npcs, dm_notes
	          FROM Sessions ORDER BY session_num`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var sessions []session
	for rows.Next() {
		var s session
		if err := rows.Scan(&s.Id, &s.SessionNum, &s.Title, &s.Chapters, &s.LevelStart, &s.LevelEnd,
			&s.Summary, &s.KeyEncounters, &s.KeyNPCs, &s.DMNotes); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		sessions = append(sessions, s)
	}
	renderTable(w, sessionsTmpl, sessions)
}

func handleLocations(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, type, danger, description, secrets FROM Locations ORDER BY name`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var locs []location
	for rows.Next() {
		var l location
		if err := rows.Scan(&l.Id, &l.Name, &l.Type_, &l.Danger, &l.Description, &l.Secrets); err != nil {
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
	ctx := context.Background()

	rows, err := db.Query(ctx,
		`SELECT id, name, chapter, COALESCE(location,''), difficulty, levelup, notes
		 FROM Encounters ORDER BY chapter, name`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var encs []encounter
	idIndex := map[int64]int{}
	for rows.Next() {
		var e encounter
		if err := rows.Scan(&e.Id, &e.Name, &e.Chapter, &e.Location, &e.Difficulty, &e.Levelup, &e.Notes); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		idIndex[e.Id] = len(encs)
		encs = append(encs, e)
	}

	// Attach monsters via join table
	mRows, err := db.Query(ctx,
		`SELECT em.encounter_id, m.name, m.cr, em.quantity
		 FROM EncounterMonsters em
		 JOIN Monsters m ON m.id = em.monster_id
		 ORDER BY em.encounter_id,
		   CASE m.cr WHEN '1/8' THEN 0.125 WHEN '1/4' THEN 0.25 WHEN '1/2' THEN 0.5
		             ELSE m.cr::numeric END NULLS LAST,
		   m.name`)
	if err != nil {
		log.Printf("encounter monsters query: %v", err)
	} else {
		defer mRows.Close()
		for mRows.Next() {
			var eid int64
			var em encounterMonster
			if err := mRows.Scan(&eid, &em.Name, &em.CR, &em.Quantity); err == nil {
				if idx, ok := idIndex[eid]; ok {
					encs[idx].Monsters = append(encs[idx].Monsters, em)
				}
			}
		}
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

func handleMonsters(w http.ResponseWriter, r *http.Request) {
	// Sort by numeric CR value, then name
	query := `SELECT id, name, type, cr, hp, hp_formula, ac, ac_desc, speed,
		str, dex, con, int_score, wis, cha,
		COALESCE(saving_throws,''), COALESCE(damage_resistances,''), COALESCE(damage_immunities,''),
		COALESCE(condition_immunities,''), COALESCE(senses,''), COALESCE(languages,''),
		COALESCE(traits,''), COALESCE(actions,''), COALESCE(legendary_actions,''), COALESCE(notes,'')
		FROM Monsters
		ORDER BY
			CASE cr WHEN '1/8' THEN 0.125 WHEN '1/4' THEN 0.25 WHEN '1/2' THEN 0.5
				ELSE cr::numeric END NULLS LAST,
			name`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var monsters []monster
	for rows.Next() {
		var m monster
		if err := rows.Scan(
			&m.Id, &m.Name, &m.Type_, &m.CR, &m.HP, &m.HPFormula,
			&m.AC, &m.ACDesc, &m.Speed,
			&m.STR, &m.DEX, &m.CON, &m.INT, &m.WIS, &m.CHA,
			&m.SavingThrows, &m.DmgResistances, &m.DmgImmunities,
			&m.CondImmunities, &m.Senses, &m.Languages,
			&m.Traits, &m.Actions, &m.LegendaryActions, &m.Notes,
		); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		monsters = append(monsters, m)
	}

	renderTable(w, monstersTmpl, monsters)
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

var sessionsTmpl = `
<div class="table-responsive">
  {{if .}}
  <table class="table table-dark table-hover table-bordered datatable w-100">
    <thead>
      <tr>
        <th>#</th>
        <th>Title</th>
        <th>Chapters</th>
        <th>Levels</th>
        <th>Summary</th>
        <th>Key Encounters</th>
        <th>Key NPCs</th>
        <th>DM Notes</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
      <tr>
        <td><strong>{{.SessionNum}}</strong></td>
        <td><strong>{{.Title}}</strong></td>
        <td><span class="badge bg-secondary">{{.Chapters}}</span></td>
        <td><span class="badge bg-info text-dark">{{.LevelStart}}{{if ne .LevelStart .LevelEnd}}→{{.LevelEnd}}{{end}}</span></td>
        <td>{{.Summary}}</td>
        <td>{{.KeyEncounters}}</td>
        <td>{{.KeyNPCs}}</td>
        <td>{{.DMNotes}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
  {{else}}<p class="empty p-3">No sessions found. Run: <code>go run ./database/seed_sessions.go</code></p>
  {{end}}
</div>`

var monstersTmpl = `
<div class="table-responsive">
  {{if .}}
  <table class="table table-dark table-hover table-bordered datatable w-100">
    <thead>
      <tr>
        <th>Name</th>
        <th>Type</th>
        <th>CR</th>
        <th>HP</th>
        <th>AC</th>
        <th>Speed</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
      <tr>
        <td><strong>{{.Name}}</strong></td>
        <td>{{.Type_}}</td>
        <td>{{.CR}}</td>
        <td>{{.HP}}{{if .HPFormula}} <small class="text-secondary">({{.HPFormula}})</small>{{end}}</td>
        <td>{{.AC}}{{if .ACDesc}} <small class="text-secondary">({{.ACDesc}})</small>{{end}}</td>
        <td>{{.Speed}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
  {{else}}<p class="empty p-3">No monsters found. Run: <code>go run ./database/seed_monsters.go</code></p>
  {{end}}
</div>`

var locationsTmpl = `
<div class="table-responsive">
  <table class="table table-dark table-hover table-bordered datatable w-100">
    <thead>
      <tr>
        <th>Name</th>
        <th>Type</th>
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
        <th>Chapter</th>
        <th>Location</th>
        <th>Difficulty</th>
        <th>Monsters</th>
        <th>Level Up</th>
        <th>Notes</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
      <tr>
        <td><strong>{{.Name}}</strong></td>
        <td>{{.Chapter}}</td>
        <td>{{.Location}}</td>
        <td class="danger" data-order="{{.Difficulty}}">{{danger .Difficulty}}</td>
        <td>
          {{range .Monsters}}
          <span class="badge bg-secondary me-1 mb-1">
            {{if ne .Quantity "1"}}{{.Quantity}}× {{end}}{{.Name}}<small class="ms-1 text-warning">CR{{.CR}}</small>
          </span>
          {{end}}
        </td>
        <td>{{if .Levelup}}<span class="badge bg-success">✓</span>{{end}}</td>
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

type encounterMonster struct {
	Name     string
	CR       string
	Quantity string
}

type encounter struct {
	Id         int64
	Name       string
	Chapter    int
	Location   string
	Difficulty int
	Monsters   []encounterMonster
	Levelup    bool
	Notes      string
}

type event struct {
	Id          int64
	Title       string
	Category    string
	Description string
}

type session struct {
	Id            int64
	SessionNum    int
	Title         string
	Chapters      string
	LevelStart    int
	LevelEnd      int
	Summary       string
	KeyEncounters string
	KeyNPCs       string
	DMNotes       string
}

type monster struct {
	Id               int64
	Name             string
	Type_            string
	CR               string
	HP               int
	HPFormula        string
	AC               int
	ACDesc           string
	Speed            string
	STR, DEX, CON    int
	INT, WIS, CHA    int
	SavingThrows     string
	DmgResistances   string
	DmgImmunities    string
	CondImmunities   string
	Senses           string
	Languages        string
	Traits           string
	Actions          string
	LegendaryActions string
	Notes            string
}
