package database

import (
	"database/sql"
	_ "embed"
	"fmt"
)

//go:embed schema.sql
var schema string

func Migrate(db *sql.DB) error {
	if _, err := db.Exec(schema); err != nil {
		return fmt.Errorf("executing schema migration: %w", err)
	}

	return nil
}
