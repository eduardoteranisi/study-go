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

func CreateSubjects(ctx context.Context, db *sql.DB, params []string, book_id int64) error {
	queries := sqlc.New(db)

	for _, param := range params {
		subject, err := queries.CreateSubject(ctx, param)
		if err != nil {
			return err
		}

		err = queries.CreateBookSubjectRelation(ctx, sqlc.CreateBookSubjectRelationParams{
			BookID:    book_id,
			SubjectID: subject.ID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateAuthors(ctx context.Context, db *sql.DB, params []string, book_id int64) error {
	queries := sqlc.New(db)

	for _, param := range params {
		author, err := queries.CreateAuthor(ctx, param)
		if err != nil {
			return err
		}

		err = queries.CreateBookAuthorRelation(ctx, sqlc.CreateBookAuthorRelationParams{
			BookID:   book_id,
			AuthorID: author.ID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateBookSubjectRelation(ctx context.Context, db *sql.DB, params sqlc.CreateBookSubjectRelationParams) error {
	queries := sqlc.New(db)

	err := queries.CreateBookSubjectRelation(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func CreateBookAuthorRelation(ctx context.Context, db *sql.DB, params sqlc.CreateBookAuthorRelationParams) error {
	queries := sqlc.New(db)

	err := queries.CreateBookAuthorRelation(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
