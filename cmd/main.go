package main

import (
	"database/sql"
	"log"

	"github.com/kidusshun/hiring_assistant/cmd/api"
	"github.com/kidusshun/hiring_assistant/config"
	"github.com/kidusshun/hiring_assistant/db"
)

func main() {
	db, err := db.NewPostgresStorage(
		config.DB.DBUser,
		config.DB.DBPassword,
		config.DB.DBAddress,
		config.DB.DBName,
	)

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	server := api.NewAPIServer(":8080", db)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB successfully connected")
}