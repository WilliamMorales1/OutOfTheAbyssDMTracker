package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const dbURL = "sqlite://oota.db?_pragma=foreign_keys(1)"

func main() {
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatalf("migrations init: %v", err)
	}
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrations down: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrations up: %v", err)
	}
	log.Println("migrations complete")
}
