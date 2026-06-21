.PHONY: build build-frontend build-backend run dev watch clean reseed

build: build-frontend build-backend

build-frontend:
	cd frontend && npm install && npm run build

build-backend:
	cd backend && go build -o oota ./cmd/oota

run: build
	cd backend && ./oota

dev: build-frontend
	cd backend && go run ./cmd/oota

watch:
	cd frontend && npm run watch
	cd backend && air

clean:
	rm -f backend/oota
	rm -rf backend/tmp
	rm -rf frontend/dist

reseed:
	cd backend && go run ./cmd/migrate
	cd backend && go run ./cmd/ingest-5etools
	cd backend && go run ./cmd/ingest-lore