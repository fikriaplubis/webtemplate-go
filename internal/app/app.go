package app

import (
	"log"

	"webtemplate/internal/config"
	"webtemplate/internal/database/postgresql"

	_ "github.com/joho/godotenv/autoload"
)

func StartApplication() {
	db, err := postgresql.SetupDb()
	if err != nil {
		log.Panicln(err.Error())
	}

	server := config.MakeServer(db)
	server.RunServer()
}
