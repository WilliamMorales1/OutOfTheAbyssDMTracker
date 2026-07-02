package main

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"regexp"
	"strconv"
	"strings"
)

// Many 5etools monster entries (e.g. "Animated Drow Statue", "Vizeran DeVir")
// aren't full stat blocks - they're a stub that says "copy this other
// monster, then apply these edits" via a "_copy" key (and optionally
// "_templates" for race templates like "Drow (Levitate)"). Without resolving
// that indirection, those monsters end up with zero HP/AC/everything because
// the fields simply aren't present in their raw JSON. This file resolves
// _copy/_templates into a fully-populated raw monster map before it's
// unmarshaled into monsterJSON.

const templateURL = "https://raw.githubusercontent.com/5etools-mirror-3/5etools-src/main/data/bestiary/template.json"

type templateFile struct {
	MonsterTemplate []struct {
		Name   string          `json:"name"`
		Source string          `json:"source"`
		Apply  json.RawMessage `json:"apply"`
	} `json:"monsterTemplate"`
}

type templateApply struct {
	Root map[string]any `json:"_root"`
	Mod  map[string]any `json:"_mod"`
}

// loadTemplates fetches the race-template definitions used by "_templates"
// references and indexes them by "Name|Source".
func loadTemplates() (map[string]templateApply, error) {
	var tf templateFile
	if err := fetchJSON(templateURL, &tf); err != nil {
		return nil, err
	}
	out := map[string]templateApply{}
	for _, t := range tf.MonsterTemplate {
		var app templateApply
		if len(t.Apply) > 0 {
			_ = json.Unmarshal(t.Apply, &app)
		}
		out[t.Name+"|"+t.Source] = app
	}
	return out, nil
}

// resolver resolves "_copy" chains against an index of every raw monster
// object seen across all ingested bestiary files.
type resolver struct {
	index     map[string]map[string]any
	templates map[string]templateApply
	cache     map[string]map[string]any
}

func newResolver(index map[string]map[string]any, templates map[string]templateApply) *resolver {
	return &resolver{index: index, templates: templates, cache: map[string]map[string]any{}}
}

func deepCopyMap(m map[string]any) map[string]any {
	b, _ := json.Marshal(m)
	var out map[string]any
	_ = json.Unmarshal(b, &out)
	return out
}

// deepCopyAny round-trips an arbitrary value through JSON. Template
// definitions are shared across every monster that references them, so
// anything spliced into a monster from a template (or from a _copy._mod
// spec) must be deep-copied first - otherwise substituteTokens mutating the
// spliced-in strings in place would corrupt the shared template for every
// other monster that uses it afterwards.
func deepCopyAny(v any) any {
	b, _ := json.Marshal(v)
	var out any
	_ = json.Unmarshal(b, &out)
	return out
}

// resolve returns the fully-merged raw monster for "Name|Source", following
// _copy and _templates chains. Results are memoized and cycles are guarded
// against by removing the copy directive before recursing.
func (r *resolver) resolve(key string) map[string]any {
	if cached, ok := r.cache[key]; ok {
		return cached
	}
	raw, ok := r.index[key]
	if !ok {
		// 5etools' own bestiary data occasionally has case mismatches
		// between a "_copy" reference and the target's actual name (e.g.
		// "Kuo-Toa" vs "Kuo-toa"). Fall back to a case-insensitive lookup.
		for k, v := range r.index {
			if strings.EqualFold(k, key) {
				raw, ok = v, true
				break
			}
		}
	}
	if !ok {
		return nil
	}
	if raw["_copy"] == nil {
		r.cache[key] = raw
		return raw
	}

	copyDirective, _ := raw["_copy"].(map[string]any)
	baseName, _ := copyDirective["name"].(string)
	baseSource, _ := copyDirective["source"].(string)
	base := r.resolve(baseName + "|" + baseSource)
	if base == nil {
		log.Printf("warning: _copy base not found for %s: %s|%s", key, baseName, baseSource)
		r.cache[key] = raw
		return raw
	}

	merged := deepCopyMap(base)
	maps.Copy(merged, raw)
	delete(merged, "_copy")
	// Prevent nested resolve() calls from re-processing this monster's own
	// copy directive if something else references it before we finish.
	r.cache[key] = merged

	if mod, ok := copyDirective["_mod"].(map[string]any); ok {
		applyModSet(merged, deepCopyAny(mod).(map[string]any))
	}

	if templRefs, ok := copyDirective["_templates"].([]any); ok {
		for _, tr := range templRefs {
			ref, _ := tr.(map[string]any)
			tname, _ := ref["name"].(string)
			tsource, _ := ref["source"].(string)
			app, ok := r.templates[tname+"|"+tsource]
			if !ok {
				continue
			}
			maps.Copy(merged, deepCopyMap(app.Root))
			if app.Mod != nil {
				applyModSet(merged, deepCopyAny(app.Mod).(map[string]any))
			}
		}
	}

	substituteTokens(merged)
	r.cache[key] = merged
	return merged
}

// applyModSet applies a "_mod" block. Keys are either a field name (operate
// on that field), "*" (apply text replacement across every string in the
// object), or "_" (whole-object operations like addSenses/addSkills).
func applyModSet(obj map[string]any, mod map[string]any) {
	for field, spec := range mod {
		ops := asOpList(spec)
		switch field {
		case "*":
			for _, op := range ops {
				applyGlobalTextOp(obj, op)
			}
		case "_":
			for _, op := range ops {
				applyWholeObjectOp(obj, op)
			}
		default:
			for _, op := range ops {
				obj[field] = applyFieldOp(obj[field], op)
			}
		}
	}
}

func asOpList(spec any) []map[string]any {
	switch v := spec.(type) {
	case map[string]any:
		return []map[string]any{v}
	case []any:
		var out []map[string]any
		for _, e := range v {
			if m, ok := e.(map[string]any); ok {
				out = append(out, m)
			}
		}
		return out
	}
	return nil
}

func applyFieldOp(field any, op map[string]any) any {
	mode, _ := op["mode"].(string)
	switch mode {
	case "replaceArr":
		replaceName, hasName := op["replace"].(string)
		items := op["items"]
		if !hasName {
			return items
		}
		arr, _ := field.([]any)
		replacements := normalizeItems(items)
		var out []any
		replaced := false
		for _, e := range arr {
			em, _ := e.(map[string]any)
			if em != nil && em["name"] == replaceName {
				out = append(out, replacements...)
				replaced = true
				continue
			}
			out = append(out, e)
		}
		if !replaced {
			out = append(out, replacements...)
		}
		return out
	case "appendArr", "appendIfNotExistsArr":
		arr, _ := field.([]any)
		items := normalizeItems(op["items"])
		for _, it := range items {
			if mode == "appendIfNotExistsArr" && containsAny(arr, it) {
				continue
			}
			arr = append(arr, it)
		}
		return arr
	case "prependArr":
		arr, _ := field.([]any)
		items := normalizeItems(op["items"])
		return append(append([]any{}, items...), arr...)
	case "removeArr":
		arr, _ := field.([]any)
		items := normalizeItems(op["items"])
		var out []any
		for _, e := range arr {
			if !containsAny(items, e) {
				out = append(out, e)
			}
		}
		return out
	case "replaceTxt":
		return replaceTxtValue(field, op)
	}
	return field
}

func normalizeItems(v any) []any {
	if arr, ok := v.([]any); ok {
		return arr
	}
	if v != nil {
		return []any{v}
	}
	return nil
}

func containsAny(arr []any, target any) bool {
	tb, _ := json.Marshal(target)
	for _, e := range arr {
		eb, _ := json.Marshal(e)
		if string(eb) == string(tb) {
			return true
		}
	}
	return false
}

// applyGlobalTextOp applies a replaceTxt over every string value reachable
// in the object (used for "*" mods like renaming "the archmage" to
// "Vizeran" throughout every trait/action/spellcasting block).
func applyGlobalTextOp(obj map[string]any, op map[string]any) {
	mode, _ := op["mode"].(string)
	if mode != "replaceTxt" {
		return
	}
	walkStrings(obj, func(s string) string {
		return replaceTxtString(s, op)
	})
}

func walkStrings(v any, fn func(string) string) {
	switch t := v.(type) {
	case map[string]any:
		for k, val := range t {
			if s, ok := val.(string); ok {
				t[k] = fn(s)
			} else {
				walkStrings(val, fn)
			}
		}
	case []any:
		for i, val := range t {
			if s, ok := val.(string); ok {
				t[i] = fn(s)
			} else {
				walkStrings(val, fn)
			}
		}
	}
}

func replaceTxtValue(field any, op map[string]any) any {
	switch t := field.(type) {
	case string:
		return replaceTxtString(t, op)
	default:
		walkStrings(field, func(s string) string { return replaceTxtString(s, op) })
		return field
	}
}

func replaceTxtString(s string, op map[string]any) string {
	replace, _ := op["replace"].(string)
	with, _ := op["with"].(string)
	if replace == "" {
		return s
	}
	flags, _ := op["flags"].(string)
	pattern := regexp.QuoteMeta(replace)
	if strings.Contains(flags, "i") {
		pattern = "(?i)" + pattern
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return s
	}
	return re.ReplaceAllString(s, with)
}

// applyWholeObjectOp handles "_" mods that touch fields outside the single
// key they're nested under, e.g. addSenses appends to the top-level "senses"
// array and addSkills merges into the top-level "skill" map.
func applyWholeObjectOp(obj map[string]any, op map[string]any) {
	mode, _ := op["mode"].(string)
	switch mode {
	case "addSenses":
		senses, _ := obj["senses"].([]any)
		if s, ok := op["senses"].(map[string]any); ok {
			typ, _ := s["type"].(string)
			rng, _ := s["range"].(float64)
			senses = append(senses, fmt.Sprintf("%s %d ft.", typ, int(rng)))
		}
		obj["senses"] = senses
	case "addSkills":
		skill, _ := obj["skill"].(map[string]any)
		if skill == nil {
			skill = map[string]any{}
		}
		if s, ok := op["skills"].(map[string]any); ok {
			for k := range s {
				if _, exists := skill[k]; !exists {
					skill[k] = "+2"
				}
			}
		}
		obj["skill"] = skill
	}
}

// --- <$token$> substitution ---

var reToken = regexp.MustCompile(`<\$([a-zA-Z_]+)\$>`)

var crProfBonus = []struct {
	min float64
	pb  int
}{
	{29, 9}, {25, 8}, {21, 7}, {17, 6}, {13, 5}, {9, 4}, {5, 3}, {0, 2},
}

func crToNumber(cr string) float64 {
	cr = strings.TrimSpace(cr)
	if strings.Contains(cr, "/") {
		parts := strings.SplitN(cr, "/", 2)
		num, _ := strconv.ParseFloat(parts[0], 64)
		den, _ := strconv.ParseFloat(parts[1], 64)
		if den != 0 {
			return num / den
		}
		return 0
	}
	n, _ := strconv.ParseFloat(cr, 64)
	return n
}

func profBonusForCR(cr string) int {
	n := crToNumber(cr)
	for _, b := range crProfBonus {
		if n >= b.min {
			return b.pb
		}
	}
	return 2
}

func abilityMod(score float64) int {
	diff := int(score) - 10
	if diff >= 0 {
		return diff / 2
	}
	return -((-diff + 1) / 2)
}

func fmtSigned(n int) string {
	if n >= 0 {
		return fmt.Sprintf("+%d", n)
	}
	return fmt.Sprintf("%d", n)
}

// substituteTokens resolves <$title_short_name$>, <$spell_dc__cha$>, etc.
// against the merged monster's own name/ability scores/CR.
func substituteTokens(obj map[string]any) {
	name, _ := obj["name"].(string)
	cr := extractCRFromAny(obj["cr"])
	pb := profBonusForCR(cr)
	abilityOf := func(key string) float64 {
		if v, ok := obj[key].(float64); ok {
			return v
		}
		return 10
	}
	walkStrings(obj, func(s string) string {
		if !strings.Contains(s, "<$") {
			return s
		}
		return reToken.ReplaceAllStringFunc(s, func(m string) string {
			sub := reToken.FindStringSubmatch(m)[1]
			switch {
			case sub == "short_name" || sub == "title_short_name":
				return name
			case strings.HasPrefix(sub, "spell_dc__"):
				ab := strings.TrimPrefix(sub, "spell_dc__")
				return strconv.Itoa(8 + pb + abilityMod(abilityOf(ab)))
			case strings.HasPrefix(sub, "to_hit__"):
				ab := strings.TrimPrefix(sub, "to_hit__")
				return fmtSigned(pb + abilityMod(abilityOf(ab)))
			case strings.HasPrefix(sub, "damage_mod__"):
				ab := strings.TrimPrefix(sub, "damage_mod__")
				return fmtSigned(abilityMod(abilityOf(ab)))
			}
			return ""
		})
	})
}

func extractCRFromAny(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case map[string]any:
		if s, ok := t["cr"].(string); ok {
			return s
		}
	}
	return "0"
}
