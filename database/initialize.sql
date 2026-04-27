-- Active: 1777184786331@@127.0.0.1@5432@oota
CREATE TABLE Locations (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text unique,
    type text,
    status text,
    danger int,
    description text,
    secrets text
);

CREATE TABLE NPCS (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    description TEXT,
    madness int,
    name text,
    disposition text,
    location text references Locations(name),
    notes text
);

CREATE TABLE Encounters (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text,
    location text references Locations(name),
    difficulty int,
    status text,
    enemies text,
    levelup boolean,
    notes text
);

CREATE TABLE Events (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title text,
    category text,
    description text
);