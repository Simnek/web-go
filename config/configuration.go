package config

import (
	"database/sql"
	"github.com/Simnek/web-go/util"
	"log"
	"time"
)

func MaxOpenConns() {
	connStr := util.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	db.SetMaxOpenConns(15)
}

func MaxIdleConns() {
	connStr := util.ConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	db.SetMaxIdleConns(5)
}

func Lifecycle() {
	connStr := util.ConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	db.SetConnMaxLifetime(100 * time.Millisecond)
	db.SetConnMaxIdleTime(5 * time.Minute)
}
