package util

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func Connect() {
	connStr := ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	log.Println("Successfully connected to database and pinged it")
}