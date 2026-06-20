# OOTA — Out of the Abyss DM Companion

A local web app for running the *Out of the Abyss* D&D 5e campaign. Tracks locations, NPCs, encounters, monsters, sessions, and campaign events. Includes a chat interface powered by a local LLM and semantic search over the campaign rulebook.

## Features

- **Campaign data browser** — sortable/searchable tables for locations, NPCs, encounters, monsters, sessions, and events
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

**3. Build the frontend**

```bash
cd frontend
npm install
npm run build
```

**4. Build and run the backend**

```bash
cd backend
go build -o oota .
./oota
```

Migrations run automatically on startup. The app listens on `http://localhost:8080` and serves the built frontend from `frontend/dist`. The binary expects to run from `backend/` (it reads `migrations/`, `images/`, and `../frontend/dist` relative to that directory).

For frontend development, run `npm run watch` in `frontend/` (recompiles on save) alongside `go run .` in `backend/`, then refresh the browser at `http://localhost:8080`.

## Database migrations

Migrations live in `backend/migrations/` and are managed by [golang-migrate](https://github.com/golang-migrate/migrate).

| Migration | Contents                                  |
| --------- | ----------------------------------------- |
| `001`   | Schema — all table definitions           |
| `002`   | Seed locations (14 Underdark locations)   |
| `003`   | Seed NPCs (25 characters)                 |
| `004`   | Seed encounters (~55 encounter entries)   |
| `005`   | Seed events (11 campaign events)          |
| `006`   | Seed sessions (16 session guides)         |
| `007`   | Seed monsters (full D&D 5e monster stats) |

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
│   ├── main.go           # HTTP server, routes, migration runner
│   ├── agent.go          # LLM chat handler, tool loop, vector embedding search
│   ├── api.go            # JSON API handlers
│   ├── db/               # sqlc-generated DB layer
│   ├── db-seeding/       # SQL schema and queries + sqlc config
│   ├── migrations/       # golang-migrate SQL migration files
│   ├── images/           # static map images, served at /images
│   └── go.mod / go.sum
└── frontend/             # TypeScript app (tsc only), built to frontend/dist
```
