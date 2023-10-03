package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
