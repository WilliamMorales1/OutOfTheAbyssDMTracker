package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

// Open opens the sqlite database at path with foreign keys enabled.
func Open(path string) (*sql.DB, error) {
	return sql.Open("sqlite", path+"?_pragma=foreign_keys(1)")
}

// RunMigrations applies all pending golang-migrate migrations in
// migrationsDir to the sqlite database at dbPath.
func RunMigrations(migrationsDir, dbPath string) error {
	dbURL := "sqlite://" + dbPath + "?_pragma=foreign_keys(1)"
	m, err := migrate.New("file://"+migrationsDir, dbURL)
	if err != nil {
		return fmt.Errorf("migrations init: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrations up: %w", err)
	}
	return nil
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBTX
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

type Monster struct {
	ID                  int64
	Name                string
	Type                sql.NullString
	Cr                  sql.NullString
	Hp                  sql.NullInt64
	HpFormula           sql.NullString
	Ac                  sql.NullInt64
	AcDesc              sql.NullString
	Speed               sql.NullString
	Str                 sql.NullInt64
	Dex                 sql.NullInt64
	Con                 sql.NullInt64
	IntScore            sql.NullInt64
	Wis                 sql.NullInt64
	Cha                 sql.NullInt64
	SavingThrows        sql.NullString
	Skills              sql.NullString
	DamageResistances   sql.NullString
	DamageImmunities    sql.NullString
	ConditionImmunities sql.NullString
	Senses              sql.NullString
	Languages           sql.NullString
	Traits              sql.NullString
	Actions             sql.NullString
	LegendaryActions    sql.NullString
	Notes               sql.NullString
	Source              sql.NullString
	Size                sql.NullString
	Alignment           sql.NullString
	PassivePerception   sql.NullInt64
	Vulnerabilities     sql.NullString
	Environment         sql.NullString
	ImageUrl            sql.NullString
	Reactions           sql.NullString
	Spellcasting        sql.NullString
	TokenUrl            sql.NullString
}

type Note struct {
	Name    string
	Content string
}

type Npc struct {
	ID          int64
	Description sql.NullString
	Madness     sql.NullInt64
	Name        sql.NullString
	Disposition sql.NullString
	Location    sql.NullString
	Notes       sql.NullString
}

type Session struct {
	ID            int64
	SessionNum    int64
	Title         string
	Chapters      sql.NullString
	LevelStart    sql.NullInt64
	LevelEnd      sql.NullInt64
	Summary       sql.NullString
	KeyEncounters sql.NullString
	KeyNpcs       sql.NullString
	Checkpoint    sql.NullString
}

const getMonster = `
SELECT id, name, type, cr, hp, hp_formula, ac, ac_desc, speed,
  str, dex, con, int_score, wis, cha,
  COALESCE(saving_throws, '')        AS saving_throws,
  COALESCE(skills, '')               AS skills,
  COALESCE(damage_resistances, '')   AS damage_resistances,
  COALESCE(damage_immunities, '')    AS damage_immunities,
  COALESCE(vulnerabilities, '')      AS vulnerabilities,
  COALESCE(condition_immunities, '') AS condition_immunities,
  COALESCE(senses, '')               AS senses,
  passive_perception,
  COALESCE(languages, '')            AS languages,
  COALESCE(traits, '')               AS traits,
  COALESCE(actions, '')              AS actions,
  COALESCE(reactions, '')            AS reactions,
  COALESCE(legendary_actions, '')    AS legendary_actions,
  COALESCE(spellcasting, '')         AS spellcasting,
  COALESCE(notes, '')                AS notes,
  COALESCE(source, '')               AS source,
  COALESCE(size, '')                 AS size,
  COALESCE(alignment, '')            AS alignment,
  COALESCE(environment, '')          AS environment,
  COALESCE(image_url, '')            AS image_url,
  COALESCE(token_url, '')            AS token_url
FROM Monsters WHERE id = ?
`

type GetMonsterRow struct {
	ID                  int64
	Name                string
	Type                sql.NullString
	Cr                  sql.NullString
	Hp                  sql.NullInt64
	HpFormula           sql.NullString
	Ac                  sql.NullInt64
	AcDesc              sql.NullString
	Speed               sql.NullString
	Str                 sql.NullInt64
	Dex                 sql.NullInt64
	Con                 sql.NullInt64
	IntScore            sql.NullInt64
	Wis                 sql.NullInt64
	Cha                 sql.NullInt64
	SavingThrows        string
	Skills              string
	DamageResistances   string
	DamageImmunities    string
	Vulnerabilities     string
	ConditionImmunities string
	Senses              string
	PassivePerception   sql.NullInt64
	Languages           string
	Traits              string
	Actions             string
	Reactions           string
	LegendaryActions    string
	Spellcasting        string
	Notes               string
	Source              string
	Size                string
	Alignment           string
	Environment         string
	ImageUrl            string
	TokenUrl            string
}

func (q *Queries) GetMonster(ctx context.Context, id int64) (GetMonsterRow, error) {
	row := q.db.QueryRowContext(ctx, getMonster, id)
	var i GetMonsterRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.Cr,
		&i.Hp,
		&i.HpFormula,
		&i.Ac,
		&i.AcDesc,
		&i.Speed,
		&i.Str,
		&i.Dex,
		&i.Con,
		&i.IntScore,
		&i.Wis,
		&i.Cha,
		&i.SavingThrows,
		&i.Skills,
		&i.DamageResistances,
		&i.DamageImmunities,
		&i.Vulnerabilities,
		&i.ConditionImmunities,
		&i.Senses,
		&i.PassivePerception,
		&i.Languages,
		&i.Traits,
		&i.Actions,
		&i.Reactions,
		&i.LegendaryActions,
		&i.Spellcasting,
		&i.Notes,
		&i.Source,
		&i.Size,
		&i.Alignment,
		&i.Environment,
		&i.ImageUrl,
		&i.TokenUrl,
	)
	return i, err
}

const getNote = `SELECT name, content FROM Notes WHERE name = ?`

func (q *Queries) GetNote(ctx context.Context, name string) (Note, error) {
	row := q.db.QueryRowContext(ctx, getNote, name)
	var i Note
	err := row.Scan(&i.Name, &i.Content)
	return i, err
}

const listMonsterStats = `SELECT id, name, ac, hp, dex FROM Monsters ORDER BY name`

type ListMonsterStatsRow struct {
	ID   int64
	Name string
	Ac   sql.NullInt64
	Hp   sql.NullInt64
	Dex  sql.NullInt64
}

func (q *Queries) ListMonsterStats(ctx context.Context) ([]ListMonsterStatsRow, error) {
	rows, err := q.db.QueryContext(ctx, listMonsterStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMonsterStatsRow
	for rows.Next() {
		var i ListMonsterStatsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Ac,
			&i.Hp,
			&i.Dex,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMonsters = `
SELECT id, name, type, cr, COALESCE(image_url, '') AS image_url
FROM Monsters
ORDER BY
  CASE cr WHEN '1/8' THEN 0.125 WHEN '1/4' THEN 0.25 WHEN '1/2' THEN 0.5
    ELSE CAST(cr AS REAL) END,
  name
`

type ListMonstersRow struct {
	ID       int64
	Name     string
	Type     sql.NullString
	Cr       sql.NullString
	ImageUrl string
}

func (q *Queries) ListMonsters(ctx context.Context) ([]ListMonstersRow, error) {
	rows, err := q.db.QueryContext(ctx, listMonsters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMonstersRow
	for rows.Next() {
		var i ListMonstersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Cr,
			&i.ImageUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listNPCs = `SELECT id, name, madness, disposition, location, notes, description FROM NPCS ORDER BY name`

type ListNPCsRow struct {
	ID          int64
	Name        sql.NullString
	Madness     sql.NullInt64
	Disposition sql.NullString
	Location    sql.NullString
	Notes       sql.NullString
	Description sql.NullString
}

func (q *Queries) ListNPCs(ctx context.Context) ([]ListNPCsRow, error) {
	rows, err := q.db.QueryContext(ctx, listNPCs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListNPCsRow
	for rows.Next() {
		var i ListNPCsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Madness,
			&i.Disposition,
			&i.Location,
			&i.Notes,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listNoteNames = `SELECT name FROM Notes ORDER BY name`

func (q *Queries) ListNoteNames(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listNoteNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listNotes = `SELECT name, content FROM Notes ORDER BY name`

func (q *Queries) ListNotes(ctx context.Context) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, listNotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Note
	for rows.Next() {
		var i Note
		if err := rows.Scan(&i.Name, &i.Content); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessions = `
SELECT id, session_num, title, chapters, level_start, level_end, summary, key_encounters, key_npcs, checkpoint
FROM Sessions ORDER BY session_num
`

func (q *Queries) ListSessions(ctx context.Context) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, listSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.SessionNum,
			&i.Title,
			&i.Chapters,
			&i.LevelStart,
			&i.LevelEnd,
			&i.Summary,
			&i.KeyEncounters,
			&i.KeyNpcs,
			&i.Checkpoint,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertMonster = `
INSERT INTO Monsters (
  name, type, cr, hp, hp_formula, ac, ac_desc, speed,
  str, dex, con, int_score, wis, cha,
  saving_throws, skills, damage_resistances, damage_immunities, vulnerabilities,
  condition_immunities, senses, passive_perception, languages,
  traits, actions, reactions, legendary_actions, spellcasting,
  source, size, alignment, environment, image_url, token_url
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?, ?
)
ON CONFLICT(name) DO UPDATE SET
  type = excluded.type, cr = excluded.cr, hp = excluded.hp, hp_formula = excluded.hp_formula,
  ac = excluded.ac, ac_desc = excluded.ac_desc, speed = excluded.speed,
  str = excluded.str, dex = excluded.dex, con = excluded.con, int_score = excluded.int_score,
  wis = excluded.wis, cha = excluded.cha,
  saving_throws = excluded.saving_throws, skills = excluded.skills,
  damage_resistances = excluded.damage_resistances, damage_immunities = excluded.damage_immunities,
  vulnerabilities = excluded.vulnerabilities, condition_immunities = excluded.condition_immunities,
  senses = excluded.senses, passive_perception = excluded.passive_perception, languages = excluded.languages,
  traits = excluded.traits, actions = excluded.actions, reactions = excluded.reactions,
  legendary_actions = excluded.legendary_actions, spellcasting = excluded.spellcasting,
  source = excluded.source, size = excluded.size, alignment = excluded.alignment,
  environment = excluded.environment, image_url = excluded.image_url, token_url = excluded.token_url
`

type UpsertMonsterParams struct {
	Name                string
	Type                sql.NullString
	Cr                  sql.NullString
	Hp                  sql.NullInt64
	HpFormula           sql.NullString
	Ac                  sql.NullInt64
	AcDesc              sql.NullString
	Speed               sql.NullString
	Str                 sql.NullInt64
	Dex                 sql.NullInt64
	Con                 sql.NullInt64
	IntScore            sql.NullInt64
	Wis                 sql.NullInt64
	Cha                 sql.NullInt64
	SavingThrows        sql.NullString
	Skills              sql.NullString
	DamageResistances   sql.NullString
	DamageImmunities    sql.NullString
	Vulnerabilities     sql.NullString
	ConditionImmunities sql.NullString
	Senses              sql.NullString
	PassivePerception   sql.NullInt64
	Languages           sql.NullString
	Traits              sql.NullString
	Actions             sql.NullString
	Reactions           sql.NullString
	LegendaryActions    sql.NullString
	Spellcasting        sql.NullString
	Source              sql.NullString
	Size                sql.NullString
	Alignment           sql.NullString
	Environment         sql.NullString
	ImageUrl            sql.NullString
	TokenUrl            sql.NullString
}

func (q *Queries) UpsertMonster(ctx context.Context, arg UpsertMonsterParams) error {
	_, err := q.db.ExecContext(ctx, upsertMonster,
		arg.Name,
		arg.Type,
		arg.Cr,
		arg.Hp,
		arg.HpFormula,
		arg.Ac,
		arg.AcDesc,
		arg.Speed,
		arg.Str,
		arg.Dex,
		arg.Con,
		arg.IntScore,
		arg.Wis,
		arg.Cha,
		arg.SavingThrows,
		arg.Skills,
		arg.DamageResistances,
		arg.DamageImmunities,
		arg.Vulnerabilities,
		arg.ConditionImmunities,
		arg.Senses,
		arg.PassivePerception,
		arg.Languages,
		arg.Traits,
		arg.Actions,
		arg.Reactions,
		arg.LegendaryActions,
		arg.Spellcasting,
		arg.Source,
		arg.Size,
		arg.Alignment,
		arg.Environment,
		arg.ImageUrl,
		arg.TokenUrl,
	)
	return err
}

const upsertNote = `
INSERT INTO Notes (name, content) VALUES (?, ?)
ON CONFLICT(name) DO UPDATE SET content = excluded.content
`

type UpsertNoteParams struct {
	Name    string
	Content string
}

func (q *Queries) UpsertNote(ctx context.Context, arg UpsertNoteParams) error {
	_, err := q.db.ExecContext(ctx, upsertNote, arg.Name, arg.Content)
	return err
}
