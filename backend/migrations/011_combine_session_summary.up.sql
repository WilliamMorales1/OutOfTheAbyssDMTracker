UPDATE Sessions
SET summary = summary
    || CASE WHEN key_encounters IS NOT NULL AND trim(key_encounters) != ''
            THEN char(10) || char(10) || 'Key Encounters: ' || key_encounters
            ELSE '' END
    || CASE WHEN key_npcs IS NOT NULL AND trim(key_npcs) != ''
            THEN char(10) || char(10) || 'Key NPCs: ' || key_npcs
            ELSE '' END;

ALTER TABLE Sessions DROP COLUMN key_encounters;
ALTER TABLE Sessions DROP COLUMN key_npcs;
