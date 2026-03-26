package main

import (
	"log"

	"github.com/icoderarely/Loopin/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(" [error] loading environment variables")
	}
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	log.Fatal(app.run())
}
