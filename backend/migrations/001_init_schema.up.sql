CREATE TABLE NPCS (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    madness     INT,
    name        TEXT,
    disposition TEXT,
    location    TEXT,
    notes       TEXT
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

CREATE TABLE Notes (
    name    TEXT PRIMARY KEY,
    content TEXT NOT NULL DEFAULT ''
);

CREATE TABLE GameMaps (
    id  TEXT PRIMARY KEY,
    img TEXT NOT NULL
);

CREATE TABLE MapMarkers (
    map_id TEXT REFERENCES GameMaps(id) ON DELETE CASCADE,
    i      INT  NOT NULL,
    x      INT  NOT NULL,
    y      INT  NOT NULL,
    title  TEXT NOT NULL,
    body   TEXT NOT NULL CHECK (length(body) <= 500),
    PRIMARY KEY (map_id, i)
);

CREATE TABLE chapter_chunks (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    chapter_title TEXT NOT NULL,
    content       TEXT NOT NULL,
    embedding     TEXT NOT NULL
);

CREATE VIRTUAL TABLE chapter_chunks_fts USING fts5(
    chapter_title,
    content,
    content='chapter_chunks',
    content_rowid='id',
    tokenize='porter unicode61'
);

CREATE TRIGGER chapter_chunks_ai AFTER INSERT ON chapter_chunks BEGIN
    INSERT INTO chapter_chunks_fts(rowid, chapter_title, content)
    VALUES (new.id, new.chapter_title, new.content);
END;

CREATE TRIGGER chapter_chunks_ad AFTER DELETE ON chapter_chunks BEGIN
    INSERT INTO chapter_chunks_fts(chapter_chunks_fts, rowid, chapter_title, content)
    VALUES ('delete', old.id, old.chapter_title, old.content);
END;

CREATE TRIGGER chapter_chunks_au AFTER UPDATE ON chapter_chunks BEGIN
    INSERT INTO chapter_chunks_fts(chapter_chunks_fts, rowid, chapter_title, content)
    VALUES ('delete', old.id, old.chapter_title, old.content);
    INSERT INTO chapter_chunks_fts(rowid, chapter_title, content)
    VALUES (new.id, new.chapter_title, new.content);
END;
