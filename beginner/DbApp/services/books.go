package services

import (
	"context"
	"database/sql"
	"strings"

	"github.com/eduardoteranisi/study-go/beginner/DbApp/database"
	"github.com/eduardoteranisi/study-go/beginner/DbApp/database/sqlc"
)

func GetBookByTitleService(ctx context.Context, db *sql.DB, title string) (*sqlc.Book, error) {
	book, err := database.GetBookByTitle(ctx, db, strings.TrimSpace(title))
	if err != nil {
		return nil, err
	}

	return book, nil
}

func GetBookByIsbnService(ctx context.Context, db *sql.DB, isbn string) (*sqlc.Book, error) {
	book, err := database.GetBookByIsbn(ctx, db, strings.TrimSpace(isbn))
	if err != nil {
		return nil, err
	}

	return book, nil
}
