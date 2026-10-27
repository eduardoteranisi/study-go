package services

import (
	"context"
	"database/sql"

	"github.com/eduardoteranisi/study-go/beginner/DbApp/database"
	"github.com/eduardoteranisi/study-go/beginner/DbApp/database/sqlc"
)

func CreateAuthorsService(ctx context.Context, db *sql.DB, authors []string, book_id int64) error {
	err := database.CreateAuthors(ctx, db, authors, book_id)
	if err != nil {
		return err
	}

	return nil
}

func GetBookByAuthorService(ctx context.Context, db *sql.DB, author string) (*sqlc.GetBookByAuthorRow, error) {
	book, err := database.GetBookByAuthor(ctx, db, author)
	if err != nil {
		return nil, err
	}

	return book, nil
}
