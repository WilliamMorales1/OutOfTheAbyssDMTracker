-- Active: 1777184786331@@127.0.0.1@5432@oota

CREATE IF NOT EXISTS TABLE Monsters (
    id          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        TEXT UNIQUE NOT NULL,
    type        TEXT,
    cr          TEXT,
    hp          INT,
    hp_formula  TEXT,
    ac          INT,
    ac_desc     TEXT,
    speed       TEXT,
    str         INT DEFAULT 10,
    dex         INT DEFAULT 10,
    con         INT DEFAULT 10,
    int_score   INT DEFAULT 10,
    wis         INT DEFAULT 10,
    cha         INT DEFAULT 10,
    saving_throws       TEXT,
    skills              TEXT,
    damage_resistances  TEXT,
    damage_immunities   TEXT,
    condition_immunities TEXT,
    senses      TEXT,
    languages   TEXT,
    traits      TEXT,
    actions     TEXT,
    legendary_actions TEXT,
    notes       TEXT
);

CREATE IF NOT EXISTS TABLE EncounterMonsters (
    encounter_id BIGINT REFERENCES Encounters(id) ON DELETE CASCADE,
    monster_id   BIGINT REFERENCES Monsters(id)   ON DELETE CASCADE,
    quantity     TEXT NOT NULL DEFAULT '1',
    PRIMARY KEY (encounter_id, monster_id)
);

ALTER TABLE Encounters DROP COLUMN enemies;

-- Run as superuser after creating tables:
-- GRANT ALL ON TABLE Monsters TO wsm52;
-- GRANT ALL ON TABLE EncounterMonsters TO wsm52;
