# OOTA — Out of the Abyss DM Companion

A local web app for running the *Out of the Abyss* D&D 5e campaign. Tracks NPCs, monsters, sessions, maps, notes, and initiative — plus a chat assistant and semantic lore search powered by a local LLM.

## Features

- **Campaign data browser** — sortable/searchable panels for NPCs, monsters, sessions, maps, and notes
- **Initiative tracker** — auto-fills combatants from the monster bestiary
- **DM chat assistant** — asks an Ollama-hosted LLM questions about the campaign; the model can query the database and search for answers
- **Lore search** — semantic vector search over the chunked *Out of the Abyss* adventure text using `nomic-embed-text-v2-moe` embeddings via Ollama
- **Auto-migrations** — SQLite schema applied automatically on startup via golang-migrate

## Requirements

- Go 1.22+
- [Ollama](https://ollama.ai) running locally on port 11434 with these models pulled (only needed for chat/lore search):
  - chat model used by the agent
  - `nomic-embed-text-v2-moe` (embeddings)

No external database setup is required — the app uses an embedded SQLite database (`backend/oota.db`), created automatically on first run.

## Setup

**1. Build and run**

From the repo root, via the `Makefile`:

```bash
make build   # builds frontend (tsc) and backend (go build -o backend/oota)
make run     # build, then run backend/oota
```

Migrations run automatically on startup. The app listens on `http://localhost:8080` and serves the built frontend from `frontend/dist`. The binary expects to run from `backend/` (it reads `migrations/`, `images/`, and `../frontend/dist` relative to that directory).

**2. Seed campaign data (optional, requires network + Ollama)**

```bash
make reseed
```

Runs migrations, then `ingest-5etools` (downloads monster stat blocks + fluff images from the 5etools data mirror into the Monsters table) and `ingest-lore` (downloads the OOTA adventure text, chunks it, embeds it, and loads it for Lore Search).

**Live reload for development:**

```bash
make watch
```

Runs `npm run watch` for the frontend and `air` for the backend (requires air: `go install github.com/air-verse/air@latest`).

Other targets: `make build-frontend`, `make build-backend`, `make dev` (frontend build + `go run`, no binary), `make clean`.

## Database migrations

Migrations live in `backend/migrations/` and are managed by [golang-migrate](https://github.com/golang-migrate/migrate) against SQLite.

| Migration | Contents                                          |
| --------- | -------------------------------------------------- |
| `001`   | Schema — all table definitions                    |
| `002`   | Seed data (NPCs, sessions, maps)                  |
| `003`   | Monster schema for 5etools-sourced bestiary data  |

To run migrations manually:

```bash
go run ./backend/cmd/migrate
```

## Tech stack

| Layer            | Tool                                                     |
| ---------------- | -------------------------------------------------------- |
| Backend language | Go                                                       |
| HTTP             | `net/http` (JSON API)                                  |
| Frontend         | TypeScript, compiled via `tsc`, no bundler/framework     |
| Database         | SQLite (`modernc.org/sqlite`)                          |
| SQL codegen      | [sqlc](https://sqlc.dev)                                    |
| Migrations       | [golang-migrate](https://github.com/golang-migrate/migrate) |
| LLM / embeddings | [Ollama](https://ollama.ai) (local)                         |

## Project structure

```
.
├── backend/
│   ├── cmd/
│   │   ├── oota/             # HTTP server: main.go, agent.go (LLM chat + tool loop), api.go (JSON API)
│   │   ├── migrate/          # standalone migration runner
│   │   ├── ingest-5etools/   # downloads monster bestiary data into the DB
│   │   └── ingest-lore/      # downloads + chunks + embeds adventure text for Lore Search
│   ├── internal/db/          # sqlc-generated DB layer
│   │   └── sqlc/             # SQL schema and queries + sqlc config
│   ├── migrations/           # golang-migrate SQL migration files (SQLite)
│   ├── images/                # static images, served at /images
│   └── go.mod / go.sum
├── frontend/
│   └── src/panels/           # chat, initiative, maps, monsters, notes, npcs, sessions, search
└── Makefile                  # build/run/watch targets for both frontend and backend
```
