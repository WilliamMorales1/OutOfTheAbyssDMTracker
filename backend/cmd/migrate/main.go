package main

import (
	"log"

	"oota/internal/db"
)

func main() {
	if err := db.ResetMigrations("migrations", "oota.db"); err != nil {
		log.Fatal(err)
	}
	log.Println("migrations complete")
}
