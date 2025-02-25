package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected to PostgreSQL")
	return db, nil
} 