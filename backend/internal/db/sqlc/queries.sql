-- name: ListNPCs :many
SELECT id, name, madness, disposition, location, notes, description FROM NPCS ORDER BY name;

-- name: ListMonsters :many
SELECT id, name, type, cr, COALESCE(image_url, '') AS image_url
FROM Monsters
ORDER BY
  CASE cr WHEN '1/8' THEN 0.125 WHEN '1/4' THEN 0.25 WHEN '1/2' THEN 0.5
    ELSE CAST(cr AS REAL) END,
  name;

-- name: ListMonsterStats :many
SELECT id, name, ac, hp, dex FROM Monsters ORDER BY name;

-- name: GetMonster :one
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
  COALESCE(image_url, '')            AS image_url
FROM Monsters WHERE id = ?;

-- name: UpsertMonster :exec
INSERT INTO Monsters (
  name, type, cr, hp, hp_formula, ac, ac_desc, speed,
  str, dex, con, int_score, wis, cha,
  saving_throws, skills, damage_resistances, damage_immunities, vulnerabilities,
  condition_immunities, senses, passive_perception, languages,
  traits, actions, reactions, legendary_actions, spellcasting,
  source, size, alignment, environment, image_url
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?
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
  environment = excluded.environment, image_url = excluded.image_url;

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
