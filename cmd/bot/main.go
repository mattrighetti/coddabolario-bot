package main

import (
	"log"
	"os"

	"coddabot"
	"coddabot/bot"
	"coddabot/database"
)

func main() {
	db, err := database.Open(os.Getenv("SQLITE_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	boot := coddabot.Bootstrap{
		DB: db,
	}

	b, err := bot.New("bot.yml", boot)
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
}
