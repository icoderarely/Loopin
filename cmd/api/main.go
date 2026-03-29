package main

import (
	"log"

	"github.com/icoderarely/Loopin/internal/env"
	"github.com/icoderarely/Loopin/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(" [error] loading environment variables")
	}

	store := store.NewPostgresStorage(nil)

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Fatal(app.run())
}
