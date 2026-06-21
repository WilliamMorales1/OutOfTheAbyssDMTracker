# OOTA — Out of the Abyss DM Companion

A local web app for running the *Out of the Abyss* D&D 5e campaign. Tracks locations, NPCs, monsters, sessions, and campaign events. Includes a chat interface powered by a local LLM and semantic search over the campaign rulebook.

## Features

- **Campaign data browser** — sortable/searchable tables for locations, NPCs, monsters, sessions, and events
- **DM chat assistant** — asks an Ollama-hosted LLM (gemma4) questions about the campaign; the model can query the database and search for answers
- **Semantic search** — vector search over chunked campaign text using `nomic-embed-text-v2-moe` embeddings via Ollama
- **Auto-migrations** — database schema and seed data applied automatically on startup via golang-migrate

## Requirements

- Go 1.22+
- PostgreSQL
- [Ollama](https://ollama.ai) running locally on port 11434 with these models pulled:
  - `gemma4` (chat)
  - `nomic-embed-text-v2-moe` (embeddings)

## Setup

**1. Create the database**

```bash
createdb oota
```

**2. Configure the connection**

The connection string is hardcoded in `backend/main.go`:

```go
const dbURL = "postgres://USER:PASSWORD@localhost/oota?sslmode=disable"
```

Update it to match your PostgreSQL credentials before building.

**3. Build and run**

From the repo root, via the `Makefile`:

```bash
make build   # builds frontend (tsc) and backend (go build -o backend/oota)
make run     # build, then run backend/oota
```

Migrations run automatically on startup. The app listens on `http://localhost:8080` and serves the built frontend from `frontend/dist`. The binary expects to run from `backend/` (it reads `migrations/`, `images/`, and `../frontend/dist` relative to that directory).

**Live reload for development:**

```bash
make watch-frontend   # tsc --watch, recompiles frontend on save
make watch-backend    # air, rebuilds + restarts backend on save (requires air: go install github.com/air-verse/air@latest)
```

Run both in separate terminals, then refresh the browser at `http://localhost:8080`.

Other targets: `make build-frontend`, `make build-backend`, `make dev` (frontend build + `go run`, no binary), `make clean`.

## Database migrations

Migrations live in `backend/migrations/` and are managed by [golang-migrate](https://github.com/golang-migrate/migrate).

| Migration | Contents                                  |
| --------- | ----------------------------------------- |
| `001`   | Schema — all table definitions           |
| `002`   | Seed locations (14 Underdark locations)   |
| `003`   | Seed NPCs (25 characters)                 |
| `005`   | Seed events (11 campaign events)          |
| `006`   | Seed sessions (16 session guides)         |
| `007`   | Seed monsters (full D&D 5e monster stats) |
| `008`   | Add checkpoint column to sessions         |
| `009`   | Seed maps                                 |
| `010`   | Drop DM notes column from sessions        |

To run migrations manually:

```bash
migrate -path backend/migrations -database "postgres://user:pass@localhost/oota?sslmode=disable" up
```

To roll back:

```bash
migrate -path backend/migrations -database "postgres://user:pass@localhost/oota?sslmode=disable" down
```

## Tech stack

| Layer            | Tool                                                     |
| ---------------- | -------------------------------------------------------- |
| Backend language | Go                                                       |
| HTTP             | `net/http` (JSON API)                                  |
| Frontend         | TypeScript, compiled via `tsc`, no bundler/framework     |
| Database driver  | `pgx/v4`                                               |
| SQL codegen      | [sqlc](https://sqlc.dev)                                    |
| Migrations       | [golang-migrate](https://github.com/golang-migrate/migrate) |
| LLM / embeddings | [Ollama](https://ollama.ai) (local)                         |

## Project structure

```
.
├── backend/
│   ├── cmd/oota/
│   │   ├── main.go       # HTTP server, routes, migration runner
│   │   ├── agent.go      # LLM chat handler, tool loop, vector embedding search
│   │   └── api.go        # JSON API handlers
│   ├── internal/db/      # sqlc-generated DB layer
│   │   └── sqlc/         # SQL schema and queries + sqlc config
│   ├── migrations/       # golang-migrate SQL migration files
│   ├── images/           # static map images, served at /images
│   ├── notes/            # session notes markdown, served via notes API
│   ├── .air.toml         # live-reload config for `make watch-backend`
│   └── go.mod / go.sum
├── frontend/             # TypeScript app (tsc only), built to frontend/dist
└── Makefile              # build/run/watch targets for both frontend and backend
```
