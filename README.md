# OOTA — Out of the Abyss DM Companion

A local web app for running the *Out of the Abyss* D&D 5e campaign. Tracks monsters, sessions, maps, notes, and initiative — plus a chat assistant and semantic lore search powered by a local LLM.

## Features

- **Sessions** — chapter-by-chapter session plan for the whole campaign, seeded from the adventure
- **Notes** — Markdown notes with a local Monaco editor and rendered preview, saved to `backend/notes/` and the database
- **Monsters & Spells** — sortable/searchable bestiary and spell list with stat blocks and token images, imported from 5etools data
- **Maps** — campaign location maps with zoom
- **Initiative tracker** — auto-fills combatants from the bestiary, tracks rounds and HP, saves preset battles
- **Soundboard** — synthesized sound effects and looping ambience via the Web Audio API, with keyboard shortcuts
- **Ask Agent** — chat with an Ollama-hosted LLM that can query the campaign database to answer questions
- **Lore Search** — semantic vector search over the chunked *Out of the Abyss* adventure text using `nomic-embed-text-v2-moe` embeddings
- **References** — quick reference tables for actions, conditions, skills, and ranges
- **Auto-migrations** — SQLite schema applied on startup by a small standard-library migration runner

## Requirements

- Go 1.26+ (pinned via `go 1.26.4` in `backend/go.mod`)
- GNU Make — runs the `Makefile` targets. Not installed by default on Windows; e.g. `winget install ezwinports.make`
- A modern web browser with native ES module support
- [Ollama](https://ollama.ai) running locally on port 11434 with these models pulled (only needed for chat/lore search):
  - any chat model with tool-calling support (picked from a dropdown in the Ask Agent panel)
  - `nomic-embed-text-v2-moe` (embeddings)

No external database setup is required — the app uses an embedded SQLite database (`backend/oota.db`), created automatically on first run.

## Setup

**1. Build and run**

From the repo root, via the `Makefile`:

```bash
make build   # validates static frontend and builds backend (go build -o backend/oota)
make run     # run backend/oota
```

Migrations run automatically on startup. The app listens on `http://localhost:8080` and serves the static frontend from `frontend/`. The binary expects to run from `backend/` (it reads `migrations/`, `images/`, and `../frontend` relative to that directory).

**Map images.** The campaign maps are copyrighted Wizards of the Coast art and are not in this repo. Drop your own copies into `backend/images/` using the filenames in `backend/migrations/002_seed_data.up.sql` (`underdark.webp`, `blingdenstone.webp`, …). The rest of the app works without them.

**2. Seed campaign data (optional, requires network + Ollama)**

```bash
make reseed
```

Runs migrations, then `ingest-5etools` (downloads monster stat blocks + fluff images from the 5etools data mirror into the Monsters table) and `ingest-lore` (downloads the OOTA adventure text, chunks it, embeds it, and loads it for Lore Search).

The frontend is plain browser JavaScript and CSS. Its runtime and production build require no npm, Node.js, package manager, or frontend build step. Markdown and the optional notes editor are vendored under `frontend/vendor/` for offline use. The optional logic test can run directly with Node.js.

Other targets: `make build-frontend`, `make build-backend`, `make test-ui` (headless Chromium UI tests), `make dev` (static frontend check + `go run`, no binary), `make clean`.

## Notes

Notes live as Markdown files in `backend/notes/` and are mirrored into the database on startup. Notes created or edited in the UI are written back to that directory. `backend/notes/*.md` is gitignored so your campaign notes stay out of version control.

## Browser UI tests

Requires Chromium installed and runs without npm:

```bash
make test-ui
```

The test starts the app, opens every panel, checks panel content, verifies local Monaco loading, opens note preview, checks core API-backed views, and compares fixed-viewport screenshots against the baselines in `backend/e2e/testdata/screenshots/`. The test runs against your local `backend/oota.db`, so data-driven panels may differ from the baselines depending on what you have ingested. Update baselines intentionally with `UPDATE_GOLDENS=1 make test-ui`.

## Database migrations

Migrations live in `backend/migrations/` and are applied against SQLite by `backend/internal/db/migrations.go`. The runner uses the compatible `schema_migrations` table, applies each migration in a transaction, and supports existing databases created by golang-migrate.

| Migration | Contents                                          |
| --------- | -------------------------------------------------- |
| `001`   | Schema — all table definitions                    |
| `002`   | Seed data (sessions, maps)                        |
| `003`   | Monster schema for 5etools-sourced bestiary data  |
| `004`   | Reference tables (actions, skills, conditions, etc.) |
| `005`   | Monster token images                              |
| `006`   | Spells                                            |
| `007`   | Monster bonus actions                             |
| `008`   | Monster lair/regional effects                     |
| `009`   | Demon lords                                       |
| `010`   | Actions and conditions reference data             |
| `011`   | Combine session summary fields                    |
| `012`   | Initiative presets                                |

Normal startup applies pending migrations. To reset and reapply every migration manually (destructive; campaign rows are reseeded):

```bash
go run ./backend/cmd/migrate
```

## Tech stack

| Layer            | Tool                                                     |
| ---------------- | -------------------------------------------------------- |
| Backend language | Go                                                       |
| HTTP             | `net/http` (JSON API)                                  |
| Frontend         | Native JavaScript modules and static CSS, no framework/build step |
| Database         | SQLite (`modernc.org/sqlite`)                          |
| Migrations       | Local runner using `database/sql` and filesystem SQL files |
| LLM / embeddings | [Ollama](https://ollama.ai) (local)                         |

## Backend dependencies

Backend keeps only two non-standard direct dependencies:

- `modernc.org/sqlite` — embedded SQLite driver. Pure Go means no CGO toolchain or system SQLite library at build time.
- `golang.org/x/image` — WebP decoding plus font/drawing helpers used by the optional `cmd/plot` diagnostic tool. Standard library image packages do not decode WebP.

All HTTP, JSON, SQL access, migrations, logging, concurrency, and embeddings client code use the Go standard library. `go mod tidy` removes unused module checksums; remaining indirect modules are required by SQLite or WebP support.

## Project structure

```
.
├── backend/
│   ├── cmd/
│   │   ├── oota/             # HTTP server: main.go, agent.go (LLM chat + tool loop), api.go (JSON API)
│   │   ├── migrate/          # standalone migration runner
│   │   ├── ingest-5etools/   # downloads monster bestiary data into the DB
│   │   ├── ingest-lore/      # downloads + chunks + embeds adventure text for Lore Search
│   │   └── plot/             # diagnostic tool for map image coordinates
│   ├── internal/db/          # hand-written DB layer (queries + models)
│   ├── migrations/           # numbered SQL migration files (SQLite)
│   ├── notes/                # your Markdown notes (gitignored)
│   ├── e2e/                  # headless Chromium UI tests and screenshot baselines
│   ├── images/               # static images, served at /images (map .webp files gitignored)
│   └── go.mod / go.sum
├── frontend/
│   ├── app/                  # native JavaScript modules and panels
│   ├── vendor/               # checked-in browser dependencies (Markdown and Monaco)
│   └── test/                 # direct Node test files (optional)
└── Makefile                  # build/run targets for both frontend and backend
```

## Disclaimer

Unofficial fan tool. *Out of the Abyss* and Dungeons & Dragons are trademarks of Wizards of the Coast. This project is not affiliated with or endorsed by Wizards of the Coast. You need your own copy of the adventure to run it.

## License

Code is released under the [MIT License](LICENSE). Vendored libraries under `frontend/vendor/` keep their own licenses. Campaign content, maps, and imported 5etools data are not covered by this license.
