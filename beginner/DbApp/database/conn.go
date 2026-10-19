package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func Close(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
