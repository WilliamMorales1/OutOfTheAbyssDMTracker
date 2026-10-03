CREATE TABLE Spells (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    name          TEXT UNIQUE NOT NULL,
    level         INT NOT NULL,
    school        TEXT,
    ritual        INT NOT NULL DEFAULT 0,
    casting_time  TEXT,
    range         TEXT,
    components    TEXT,
    duration      TEXT,
    concentration INT NOT NULL DEFAULT 0,
    classes       TEXT,
    description   TEXT,
    higher_level  TEXT,
    source        TEXT
);
