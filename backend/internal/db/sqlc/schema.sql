CREATE TABLE Locations (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT UNIQUE,
    type        TEXT,
    status      TEXT,
    danger      INT,
    description TEXT,
    secrets     TEXT
);

CREATE TABLE NPCS (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    madness     INT,
    name        TEXT,
    disposition TEXT,
    location    TEXT REFERENCES Locations(name),
    notes       TEXT
);

CREATE TABLE Encounters (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT,
    location    TEXT REFERENCES Locations(name),
    difficulty  INT,
    status      TEXT,
    enemies     TEXT,
    chapter     INTEGER,
    levelup     BOOLEAN,
    notes       TEXT
);

CREATE TABLE Events (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT,
    category    TEXT,
    description TEXT
);

CREATE TABLE Monsters (
    id                   INTEGER PRIMARY KEY AUTOINCREMENT,
    name                 TEXT UNIQUE NOT NULL,
    type                 TEXT,
    cr                   TEXT,
    hp                   INT,
    hp_formula           TEXT,
    ac                   INT,
    ac_desc              TEXT,
    speed                TEXT,
    str                  INT DEFAULT 10,
    dex                  INT DEFAULT 10,
    con                  INT DEFAULT 10,
    int_score            INT DEFAULT 10,
    wis                  INT DEFAULT 10,
    cha                  INT DEFAULT 10,
    saving_throws        TEXT,
    skills               TEXT,
    damage_resistances   TEXT,
    damage_immunities    TEXT,
    condition_immunities TEXT,
    senses               TEXT,
    languages            TEXT,
    traits               TEXT,
    actions              TEXT,
    legendary_actions    TEXT,
    notes                TEXT
);

CREATE TABLE EncounterMonsters (
    encounter_id INTEGER REFERENCES Encounters(id) ON DELETE CASCADE,
    monster_id   INTEGER REFERENCES Monsters(id)   ON DELETE CASCADE,
    quantity     TEXT NOT NULL DEFAULT '1',
    PRIMARY KEY (encounter_id, monster_id)
);

CREATE TABLE Sessions (
    id             INTEGER PRIMARY KEY AUTOINCREMENT,
    session_num    INTEGER UNIQUE NOT NULL,
    title          TEXT NOT NULL,
    chapters       TEXT,
    level_start    INTEGER,
    level_end      INTEGER,
    summary        TEXT,
    key_encounters TEXT,
    key_npcs       TEXT,
    checkpoint     TEXT
);
