// ingest-5etools downloads monster stat blocks (and their fluff images) from
// the 5etools data mirror and loads them into the Monsters table, replacing
// the old hand-written seed data with the full published bestiary.
//
// Run manually (requires network access):
//
//	go run ./cmd/ingest-5etools
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"oota/internal/db"

	_ "modernc.org/sqlite"
)

const dataBaseURL = "https://raw.githubusercontent.com/5etools-mirror-3/5etools-src/main/data/bestiary/"
const imgBaseURL = "https://5e.tools/img/"

// sources lists the bestiary files to ingest: the core Monster Manual plus
// every book that contributes monsters to Out of the Abyss (demons, drow,
// Underdark creatures).
var sources = []string{"mm", "oota", "mtf", "vgm"}

type entryBlock struct {
	Name    string `json:"name"`
	Entries []any  `json:"entries"`
}

type spellcastingBlock struct {
	Name          string         `json:"name"`
	HeaderEntries []string       `json:"headerEntries"`
	FooterEntries []string       `json:"footerEntries"`
	Spells        map[string]any `json:"spells"`
	Will          []string       `json:"will"`
	Daily         map[string]any `json:"daily"`
}

type monsterJSON struct {
	Name      string            `json:"name"`
	Source    string            `json:"source"`
	Size      []string          `json:"size"`
	Type      json.RawMessage   `json:"type"`
	Alignment []json.RawMessage `json:"alignment"`
	AC        []json.RawMessage `json:"ac"`
	HP        *struct {
		Average int    `json:"average"`
		Formula string `json:"formula"`
	} `json:"hp"`
	Speed           map[string]json.RawMessage `json:"speed"`
	Str             *int                       `json:"str"`
	Dex             *int                       `json:"dex"`
	Con             *int                       `json:"con"`
	Int             *int                       `json:"int"`
	Wis             *int                       `json:"wis"`
	Cha             *int                       `json:"cha"`
	Save            map[string]string          `json:"save"`
	Skill           map[string]json.RawMessage `json:"skill"`
	Senses          []string                   `json:"senses"`
	Passive         *int                       `json:"passive"`
	Immune          []json.RawMessage          `json:"immune"`
	Resist          []json.RawMessage          `json:"resist"`
	Vulnerable      []json.RawMessage          `json:"vulnerable"`
	ConditionImmune []json.RawMessage          `json:"conditionImmune"`
	Languages       []string                   `json:"languages"`
	CR              json.RawMessage            `json:"cr"`
	Trait           []entryBlock               `json:"trait"`
	Action          []entryBlock               `json:"action"`
	Reaction        []entryBlock               `json:"reaction"`
	Legendary       []entryBlock               `json:"legendary"`
	Spellcasting    []spellcastingBlock        `json:"spellcasting"`
	Environment     []string                   `json:"environment"`
}

type bestiaryFile struct {
	Monster []monsterJSON `json:"monster"`
}

type bestiaryFileRaw struct {
	Monster []map[string]any `json:"monster"`
}

type fluffImage struct {
	Type string `json:"type"`
	Href struct {
		Type string `json:"type"`
		Path string `json:"path"`
	} `json:"href"`
}

type monsterFluffJSON struct {
	Name   string       `json:"name"`
	Source string       `json:"source"`
	Images []fluffImage `json:"images"`
}

type fluffFile struct {
	MonsterFluff []monsterFluffJSON `json:"monsterFluff"`
}

// --- 5etools {@tag ...} markup cleanup ---

var (
	reAtkTag      = regexp.MustCompile(`\{@atk ([a-z,]+)\}`)
	reHitTag      = regexp.MustCompile(`\{@h\}`)
	reHitBonusTag = regexp.MustCompile(`\{@hit (-?\d+)\}`)
	reDCTag       = regexp.MustCompile(`\{@dc (\d+)\}`)
	reRechargeTag = regexp.MustCompile(`\{@recharge ?(\d?)\}`)
	reGenericTag  = regexp.MustCompile(`\{@\w+ ([^}|]+)(\|[^}]*)?\}`)
	reWS          = regexp.MustCompile(`\s+`)
)

var atkLabels = map[string]string{
	"m":     "Melee Attack:",
	"mw":    "Melee Weapon Attack:",
	"rw":    "Ranged Weapon Attack:",
	"mw,rw": "Melee or Ranged Weapon Attack:",
	"rw,mw": "Melee or Ranged Weapon Attack:",
	"ms":    "Melee Spell Attack:",
	"rs":    "Ranged Spell Attack:",
}

func cleanText(s string) string {
	s = reAtkTag.ReplaceAllStringFunc(s, func(m string) string {
		key := reAtkTag.FindStringSubmatch(m)[1]
		if label, ok := atkLabels[key]; ok {
			return label
		}
		return "Attack:"
	})
	s = reHitTag.ReplaceAllString(s, "Hit: ")
	s = reHitBonusTag.ReplaceAllString(s, "+$1")
	s = reDCTag.ReplaceAllString(s, "DC $1")
	s = reRechargeTag.ReplaceAllStringFunc(s, func(m string) string {
		sub := reRechargeTag.FindStringSubmatch(m)
		if sub[1] == "" {
			return "(Recharge 6)"
		}
		return "(Recharge " + sub[1] + "-6)"
	})
	s = reGenericTag.ReplaceAllString(s, "$1")
	s = reWS.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// flattenEntries walks an arbitrary 5etools entries tree into plain-text
// paragraphs, joined with newlines.
func flattenEntries(entries []any) string {
	var out []string
	var walk func(e any)
	walk = func(e any) {
		switch v := e.(type) {
		case string:
			if t := cleanText(v); t != "" {
				out = append(out, t)
			}
		case map[string]any:
			if name, ok := v["name"].(string); ok {
				if items, ok := v["entries"].([]any); ok {
					sub := flattenEntries(items)
					if sub != "" {
						out = append(out, name+": "+sub)
					}
					return
				}
			}
			if items, ok := v["items"].([]any); ok {
				for _, it := range items {
					walk(it)
				}
				return
			}
			if items, ok := v["entries"].([]any); ok {
				for _, it := range items {
					walk(it)
				}
			}
		case []any:
			for _, it := range v {
				walk(it)
			}
		}
	}
	for _, e := range entries {
		walk(e)
	}
	return strings.Join(out, "\n")
}

func entryBlocksToJSON(blocks []entryBlock) string {
	type out struct {
		Name string `json:"name"`
		Text string `json:"text"`
	}
	var res []out
	for _, b := range blocks {
		text := flattenEntries(b.Entries)
		if text == "" {
			continue
		}
		res = append(res, out{Name: cleanText(b.Name), Text: text})
	}
	if len(res) == 0 {
		return ""
	}
	j, _ := json.Marshal(res)
	return string(j)
}

func spellcastingToJSON(blocks []spellcastingBlock) string {
	type out struct {
		Name string `json:"name"`
		Text string `json:"text"`
	}
	var res []out
	for _, b := range blocks {
		var parts []string
		for _, h := range b.HeaderEntries {
			if t := cleanText(h); t != "" {
				parts = append(parts, t)
			}
		}
		// Spell levels/groups, sorted by key (0=cantrips first).
		if len(b.Spells) > 0 {
			keys := make([]string, 0, len(b.Spells))
			for k := range b.Spells {
				keys = append(keys, k)
			}
			sortStrings(keys)
			for _, k := range keys {
				group, _ := b.Spells[k].(map[string]any)
				if group == nil {
					continue
				}
				spells, _ := group["spells"].([]any)
				var names []string
				for _, sp := range spells {
					if s, ok := sp.(string); ok {
						names = append(names, cleanText(s))
					}
				}
				if len(names) == 0 {
					continue
				}
				label := "Cantrips"
				if k != "0" {
					label = ordinalLevel(k) + " Level"
					if slots, ok := group["slots"].(float64); ok {
						label = fmt.Sprintf("%s (%d slots)", label, int(slots))
					}
				}
				parts = append(parts, label+": "+strings.Join(names, ", "))
			}
		}
		if len(b.Will) > 0 {
			var names []string
			for _, s := range b.Will {
				names = append(names, cleanText(s))
			}
			parts = append(parts, "At will: "+strings.Join(names, ", "))
		}
		for _, f := range b.FooterEntries {
			if t := cleanText(f); t != "" {
				parts = append(parts, t)
			}
		}
		if len(parts) == 0 {
			continue
		}
		name := b.Name
		if name == "" {
			name = "Spellcasting"
		}
		res = append(res, out{Name: name, Text: strings.Join(parts, "\n")})
	}
	if len(res) == 0 {
		return ""
	}
	j, _ := json.Marshal(res)
	return string(j)
}

func sortStrings(s []string) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}

func ordinalLevel(k string) string {
	switch k {
	case "1":
		return "1st"
	case "2":
		return "2nd"
	case "3":
		return "3rd"
	default:
		return k + "th"
	}
}

// --- field extraction ---

func extractType(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		return s
	}
	var obj struct {
		Type string   `json:"type"`
		Tags []string `json:"tags"`
	}
	if err := json.Unmarshal(raw, &obj); err == nil {
		if len(obj.Tags) > 0 {
			return obj.Type + " (" + strings.Join(obj.Tags, ", ") + ")"
		}
		return obj.Type
	}
	return ""
}

func extractAlignment(raw []json.RawMessage) string {
	var parts []string
	for _, r := range raw {
		var s string
		if err := json.Unmarshal(r, &s); err == nil {
			parts = append(parts, s)
			continue
		}
		var obj struct {
			Alignment []string `json:"alignment"`
			Chance    float64  `json:"chance"`
		}
		if err := json.Unmarshal(r, &obj); err == nil {
			parts = append(parts, strings.Join(obj.Alignment, " "))
		}
	}
	return strings.Join(parts, " ")
}

func extractAC(raw []json.RawMessage) (int64, string) {
	if len(raw) == 0 {
		return 0, ""
	}
	var n float64
	if err := json.Unmarshal(raw[0], &n); err == nil {
		return int64(n), ""
	}
	var obj struct {
		AC        float64  `json:"ac"`
		From      []string `json:"from"`
		Condition string   `json:"condition"`
	}
	if err := json.Unmarshal(raw[0], &obj); err == nil {
		desc := strings.Join(obj.From, ", ")
		if obj.Condition != "" {
			d := cleanText(obj.Condition)
			if desc != "" {
				desc += " " + d
			} else {
				desc = d
			}
		}
		return int64(obj.AC), cleanText(desc)
	}
	return 0, ""
}

func extractSpeed(m map[string]json.RawMessage) string {
	order := []string{"walk", "fly", "swim", "climb", "burrow"}
	var parts []string
	for _, k := range order {
		raw, ok := m[k]
		if !ok {
			continue
		}
		var n float64
		if err := json.Unmarshal(raw, &n); err == nil {
			label := k
			if k == "walk" {
				parts = append(parts, fmt.Sprintf("%d ft.", int(n)))
			} else {
				parts = append(parts, fmt.Sprintf("%s %d ft.", label, int(n)))
			}
			continue
		}
		var obj struct {
			Number    float64 `json:"number"`
			Condition string  `json:"condition"`
		}
		if err := json.Unmarshal(raw, &obj); err == nil {
			seg := fmt.Sprintf("%s %d ft.", k, int(obj.Number))
			if k == "walk" {
				seg = fmt.Sprintf("%d ft.", int(obj.Number))
			}
			if obj.Condition != "" {
				seg += " " + cleanText(obj.Condition)
			}
			parts = append(parts, seg)
		}
	}
	if hover, ok := m["canHover"]; ok {
		var b bool
		if json.Unmarshal(hover, &b) == nil && b {
			parts = append(parts, "(hover)")
		}
	}
	return strings.Join(parts, ", ")
}

func extractCR(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		return s
	}
	var obj struct {
		CR string `json:"cr"`
	}
	if err := json.Unmarshal(raw, &obj); err == nil {
		return obj.CR
	}
	return ""
}

// extractDamageList flattens 5etools' damage list shape (strings or
// {resist:[...],note} objects) into a single comma-separated string.
func extractDamageList(raw []json.RawMessage) string {
	var parts []string
	for _, r := range raw {
		var s string
		if err := json.Unmarshal(r, &s); err == nil {
			parts = append(parts, cleanText(s))
			continue
		}
		var obj struct {
			Resist          []string `json:"resist"`
			Immune          []string `json:"immune"`
			Vulnerable      []string `json:"vulnerable"`
			ConditionImmune []string `json:"conditionImmune"`
			Note            string   `json:"note"`
		}
		if err := json.Unmarshal(r, &obj); err == nil {
			combined := append(append(append(obj.Resist, obj.Immune...), obj.Vulnerable...), obj.ConditionImmune...)
			s := strings.Join(combined, ", ")
			if obj.Note != "" {
				s += " " + cleanText(obj.Note)
			}
			parts = append(parts, strings.TrimSpace(s))
		}
	}
	return strings.Join(parts, "; ")
}

func savesToString(m map[string]string) string {
	var parts []string
	for _, k := range []string{"str", "dex", "con", "int", "wis", "cha"} {
		if v, ok := m[k]; ok {
			parts = append(parts, strings.ToUpper(k[:1])+k[1:]+" "+v)
		}
	}
	return strings.Join(parts, ", ")
}

func skillsToString(m map[string]json.RawMessage) string {
	var parts []string
	for k, raw := range m {
		var v string
		if err := json.Unmarshal(raw, &v); err != nil {
			// "other" holds conditional alternative skill sets we don't
			// surface in the flattened display - skip rather than guess.
			continue
		}
		title := strings.ToUpper(k[:1]) + k[1:]
		parts = append(parts, title+" "+v)
	}
	sortStrings(parts)
	return strings.Join(parts, ", ")
}

func intVal(p *int) int64 {
	if p == nil {
		return 10
	}
	return int64(*p)
}

func imageURL(fluff map[string]monsterFluffJSON, name, source string) string {
	f, ok := fluff[name+"|"+source]
	if !ok || len(f.Images) == 0 {
		return ""
	}
	return imgBaseURL + f.Images[0].Href.Path
}

// tokenURL builds the 5etools token image URL for a monster. Tokens aren't
// listed in the fluff JSON - the site derives them by convention from the
// monster's name/source, so we do the same (mirroring 5etools'
// Renderer.monster.getTokenUrl).
func tokenURL(name, source string) string {
	return imgBaseURL + "bestiary/tokens/" + source + "/" + url.PathEscape(name) + ".webp"
}

func fetchJSON(url string, v any) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fetch %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("fetch %s: status %s", url, resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

func main() {
	dbPath := flag.String("db", "oota.db", "path to the sqlite database")
	flag.Parse()

	conn, err := db.Open(*dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer conn.Close()
	q := db.New(conn)

	fluff := map[string]monsterFluffJSON{}
	rawIndex := map[string]map[string]any{}
	var rawOrder []string

	for _, src := range sources {
		var bf bestiaryFileRaw
		log.Printf("fetching bestiary-%s.json", src)
		if err := fetchJSON(dataBaseURL+"bestiary-"+src+".json", &bf); err != nil {
			log.Fatalf("%v", err)
		}
		for _, m := range bf.Monster {
			name, _ := m["name"].(string)
			source, _ := m["source"].(string)
			key := name + "|" + source
			if _, exists := rawIndex[key]; !exists {
				rawOrder = append(rawOrder, key)
			}
			rawIndex[key] = m
		}

		var ff fluffFile
		log.Printf("fetching fluff-bestiary-%s.json", src)
		if err := fetchJSON(dataBaseURL+"fluff-bestiary-"+src+".json", &ff); err != nil {
			log.Printf("warning: no fluff for %s: %v", src, err)
			continue
		}
		for _, mf := range ff.MonsterFluff {
			fluff[mf.Name+"|"+mf.Source] = mf
		}
	}

	log.Printf("fetching template.json")
	templates, err := loadTemplates()
	if err != nil {
		log.Printf("warning: failed to load monster templates, _copy monsters using them will be incomplete: %v", err)
		templates = map[string]templateApply{}
	}

	res := newResolver(rawIndex, templates)
	var monsters []monsterJSON
	for _, key := range rawOrder {
		merged := res.resolve(key)
		b, err := json.Marshal(merged)
		if err != nil {
			log.Fatalf("marshal resolved monster %q: %v", key, err)
		}
		var mj monsterJSON
		if err := json.Unmarshal(b, &mj); err != nil {
			log.Fatalf("unmarshal resolved monster %q: %v", key, err)
		}
		monsters = append(monsters, mj)
	}
	log.Printf("parsed %d monsters from %d sources", len(monsters), len(sources))

	// Build params for every monster in parallel - the Go scheduler decides
	// how many goroutines actually run concurrently (one per monster, no
	// artificial semaphore); a single writer goroutine serializes the
	// sqlite inserts since sqlite only supports one writer at a time.
	type result struct {
		params db.UpsertMonsterParams
	}
	resultCh := make(chan result)
	var wg sync.WaitGroup
	seen := map[string]bool{}
	for _, mj := range monsters {
		key := mj.Name + "|" + mj.Source
		if seen[key] {
			continue
		}
		seen[key] = true
		wg.Add(1)
		go func(mj monsterJSON) {
			defer wg.Done()
			ac, acDesc := extractAC(mj.AC)
			hp, hpFormula := int64(0), ""
			if mj.HP != nil {
				hp, hpFormula = int64(mj.HP.Average), mj.HP.Formula
			}
			passive := int64(10)
			if mj.Passive != nil {
				passive = int64(*mj.Passive)
			}
			resultCh <- result{params: db.UpsertMonsterParams{
				Name:                mj.Name,
				Type:                sql.NullString{String: extractType(mj.Type), Valid: true},
				Cr:                  sql.NullString{String: extractCR(mj.CR), Valid: true},
				Hp:                  sql.NullInt64{Int64: hp, Valid: true},
				HpFormula:           sql.NullString{String: hpFormula, Valid: true},
				Ac:                  sql.NullInt64{Int64: ac, Valid: true},
				AcDesc:              sql.NullString{String: acDesc, Valid: true},
				Speed:               sql.NullString{String: extractSpeed(mj.Speed), Valid: true},
				Str:                 sql.NullInt64{Int64: intVal(mj.Str), Valid: true},
				Dex:                 sql.NullInt64{Int64: intVal(mj.Dex), Valid: true},
				Con:                 sql.NullInt64{Int64: intVal(mj.Con), Valid: true},
				IntScore:            sql.NullInt64{Int64: intVal(mj.Int), Valid: true},
				Wis:                 sql.NullInt64{Int64: intVal(mj.Wis), Valid: true},
				Cha:                 sql.NullInt64{Int64: intVal(mj.Cha), Valid: true},
				SavingThrows:        sql.NullString{String: savesToString(mj.Save), Valid: true},
				Skills:              sql.NullString{String: skillsToString(mj.Skill), Valid: true},
				DamageResistances:   sql.NullString{String: extractDamageList(mj.Resist), Valid: true},
				DamageImmunities:    sql.NullString{String: extractDamageList(mj.Immune), Valid: true},
				Vulnerabilities:     sql.NullString{String: extractDamageList(mj.Vulnerable), Valid: true},
				ConditionImmunities: sql.NullString{String: extractDamageList(mj.ConditionImmune), Valid: true},
				Senses:              sql.NullString{String: strings.Join(mj.Senses, ", "), Valid: true},
				PassivePerception:   sql.NullInt64{Int64: passive, Valid: true},
				Languages:           sql.NullString{String: strings.Join(mj.Languages, ", "), Valid: true},
				Traits:              sql.NullString{String: entryBlocksToJSON(mj.Trait), Valid: true},
				Actions:             sql.NullString{String: entryBlocksToJSON(mj.Action), Valid: true},
				Reactions:           sql.NullString{String: entryBlocksToJSON(mj.Reaction), Valid: true},
				LegendaryActions:    sql.NullString{String: entryBlocksToJSON(mj.Legendary), Valid: true},
				Spellcasting:        sql.NullString{String: spellcastingToJSON(mj.Spellcasting), Valid: true},
				Source:              sql.NullString{String: mj.Source, Valid: true},
				Size:                sql.NullString{String: strings.Join(mj.Size, ""), Valid: true},
				Alignment:           sql.NullString{String: extractAlignment(mj.Alignment), Valid: true},
				Environment:         sql.NullString{String: strings.Join(mj.Environment, ", "), Valid: true},
				ImageUrl:            sql.NullString{String: imageURL(fluff, mj.Name, mj.Source), Valid: true},
				TokenUrl:            sql.NullString{String: tokenURL(mj.Name, mj.Source), Valid: true},
			}}
		}(mj)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	ctx := context.Background()
	total := 0
	for res := range resultCh {
		if err := q.UpsertMonster(ctx, res.params); err != nil {
			log.Fatalf("upsert %q: %v", res.params.Name, err)
		}
		total++
		if total%50 == 0 {
			log.Printf("progress: %d/%d", total, len(seen))
		}
	}
	log.Printf("done: %d monsters ingested", total)
}
