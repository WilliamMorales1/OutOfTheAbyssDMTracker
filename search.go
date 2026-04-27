package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

const indexPath = "oota.bleve"

var bIdx bleve.Index

func openIndex() error {
	var err error
	bIdx, err = bleve.Open(indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		mapping := bleve.NewIndexMapping()
		bIdx, err = bleve.New(indexPath, mapping)
	}
	return err
}

type doc struct {
	Table       string `json:"table"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Danger      int    `json:"danger"`
	Description string `json:"description"`
	Secrets     string `json:"secrets"`
	Madness     int    `json:"madness"`
	Disposition string `json:"disposition"`
	Location    string `json:"location"`
	Notes       string `json:"notes"`
	Difficulty  int    `json:"difficulty"`
	Enemies     string `json:"enemies"`
	Category    string `json:"category"`
}

func indexAll() error {
	if bIdx == nil {
		if err := openIndex(); err != nil {
			return err
		}
	}

	batch := bIdx.NewBatch()
	ctx := context.Background()

	rows, err := db.Query(ctx, "SELECT id, name, type, status, danger, description, secrets FROM Locations")
	if err != nil {
		return err
	}
	for rows.Next() {
		var l location
		rows.Scan(&l.Id, &l.Name, &l.Type_, &l.Status, &l.Danger, &l.Description, &l.Secrets)
		batch.Index(fmt.Sprintf("loc-%d", l.Id), doc{
			Table: "locations", Name: l.Name, Type: l.Type_, Status: l.Status,
			Danger: l.Danger, Description: l.Description, Secrets: l.Secrets,
		})
	}
	rows.Close()

	rows, err = db.Query(ctx, "SELECT id, name, madness, disposition, location, notes, description FROM NPCS")
	if err != nil {
		return err
	}
	for rows.Next() {
		var n npc
		rows.Scan(&n.Id, &n.Name, &n.Madness, &n.Disposition, &n.Location, &n.Notes, &n.Description)
		batch.Index(fmt.Sprintf("npc-%d", n.Id), doc{
			Table: "npcs", Name: n.Name, Madness: n.Madness, Disposition: n.Disposition,
			Location: n.Location, Notes: n.Notes, Description: n.Description,
		})
	}
	rows.Close()

	rows, err = db.Query(ctx, "SELECT id, name, location, difficulty, status, enemies, levelup, notes FROM Encounters")
	if err != nil {
		return err
	}
	for rows.Next() {
		var e encounter
		rows.Scan(&e.Id, &e.Name, &e.Location, &e.Difficulty, &e.Status, &e.Enemies, &e.Levelup, &e.Notes)
		batch.Index(fmt.Sprintf("enc-%d", e.Id), doc{
			Table: "encounters", Name: e.Name, Location: e.Location, Difficulty: e.Difficulty,
			Status: e.Status, Enemies: e.Enemies, Notes: e.Notes,
		})
	}
	rows.Close()

	rows, err = db.Query(ctx, "SELECT id, title, category, description FROM Events")
	if err != nil {
		return err
	}
	for rows.Next() {
		var ev event
		rows.Scan(&ev.Id, &ev.Title, &ev.Category, &ev.Description)
		batch.Index(fmt.Sprintf("evt-%d", ev.Id), doc{
			Table: "events", Title: ev.Title, Category: ev.Category, Description: ev.Description,
		})
	}
	rows.Close()

	return bIdx.Batch(batch)
}

func searchBleve(query, table string) string {
	if bIdx == nil {
		return "Search index not ready."
	}

	q := bleve.NewFuzzyQuery(query)
	req := bleve.NewSearchRequest(q)
	req.Size = 10
	req.Fields = []string{"*"}

	if table != "" && table != "all" {
		tq := bleve.NewTermQuery(strings.ToLower(table))
		tq.SetField("table")
		req.Query = bleve.NewConjunctionQuery(tq, q)
	}

	res, err := bIdx.Search(req)
	if err != nil {
		return "Search error: " + err.Error()
	}
	if res.Total == 0 {
		return "No results found."
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Found %d result(s):\n\n", res.Total))
	for i, hit := range res.Hits {
		tbl, _ := hit.Fields["table"].(string)
		name, _ := hit.Fields["name"].(string)
		title, _ := hit.Fields["title"].(string)
		label := name
		if label == "" {
			label = title
		}
		sb.WriteString(fmt.Sprintf("[%d] (%s) %s", i+1, strings.ToUpper(tbl), label))
		for _, k := range []string{"status", "disposition", "location", "category", "description", "notes", "secrets", "enemies"} {
			if v, ok := hit.Fields[k].(string); ok && v != "" {
				sb.WriteString(fmt.Sprintf(" %s=%s", k, v))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
