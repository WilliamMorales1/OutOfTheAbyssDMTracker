-- name: ListNPCs :many
SELECT id, name, madness, disposition, location, notes, description FROM NPCS ORDER BY name;

-- name: ListEncounters :many
SELECT id, name, chapter, COALESCE(location, '') AS location, difficulty, levelup, notes
FROM Encounters ORDER BY chapter, name;

-- name: ListEncounterMonsters :many
SELECT em.encounter_id, m.name, m.cr, em.quantity
FROM EncounterMonsters em
JOIN Monsters m ON m.id = em.monster_id
ORDER BY em.encounter_id,
  CASE m.cr WHEN '1/8' THEN 0.125 WHEN '1/4' THEN 0.25 WHEN '1/2' THEN 0.5
            ELSE CAST(m.cr AS REAL) END,
  m.name;

-- name: ListMonsters :many
SELECT id, name, type, cr, hp, hp_formula, ac, ac_desc, speed,
  str, dex, con, int_score, wis, cha,
  COALESCE(saving_throws, '')        AS saving_throws,
  COALESCE(damage_resistances, '')   AS damage_resistances,
  COALESCE(damage_immunities, '')    AS damage_immunities,
  COALESCE(condition_immunities, '') AS condition_immunities,
  COALESCE(senses, '')               AS senses,
  COALESCE(languages, '')            AS languages,
  COALESCE(traits, '')               AS traits,
  COALESCE(actions, '')              AS actions,
  COALESCE(legendary_actions, '')    AS legendary_actions,
  COALESCE(notes, '')                AS notes
FROM Monsters
ORDER BY
  CASE cr WHEN '1/8' THEN 0.125 WHEN '1/4' THEN 0.25 WHEN '1/2' THEN 0.5
    ELSE CAST(cr AS REAL) END,
  name;

-- name: ListSessions :many
SELECT id, session_num, title, chapters, level_start, level_end, summary, key_encounters, key_npcs, checkpoint
FROM Sessions ORDER BY session_num;

-- name: ListNoteNames :many
SELECT name FROM Notes ORDER BY name;

-- name: ListNotes :many
SELECT name, content FROM Notes ORDER BY name;

-- name: GetNote :one
SELECT name, content FROM Notes WHERE name = ?;

-- name: UpsertNote :exec
INSERT INTO Notes (name, content) VALUES (?, ?)
ON CONFLICT(name) DO UPDATE SET content = excluded.content;
