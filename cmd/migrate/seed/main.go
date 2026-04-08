package main

import (
	"log"

	"github.com/icoderarely/Loopin/internal/db"
	"github.com/icoderarely/Loopin/internal/env"
	"github.com/icoderarely/Loopin/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:pass@localhost/loopin?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewPostgresStorage(conn)

	db.Seed(store, conn)
}
