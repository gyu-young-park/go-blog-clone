package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
}

func StartDB(driverString string, dsn string) *Store {
	db, err := sql.Open(driverString, dsn)
	if err != nil {
		log.Fatal("failed to open database", err.Error())
	}

	return &Store{
		db: db,
	}
}
