package database

import (
	"context"
	"database/sql"

	"github.com/eduardoteranisi/study-go/beginner/DbApp/database/sqlc"
)

func CreateBook(ctx context.Context, db *sql.DB, params sqlc.CreateBookParams) (*sqlc.Book, error) {
	queries := sqlc.New(db)

	book, err := queries.CreateBook(ctx, params)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
