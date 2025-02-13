package main

import (
	"database/sql"
	"log"

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
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB successfully connected")
}