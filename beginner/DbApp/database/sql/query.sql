-- name: CreateBook :one
INSERT INTO books (isbn, title, subtitle, publisher, year, format, page_count) 
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: CreateSubject :one
INSERT INTO subjects (name)
VALUES (?)
RETURNING *;

-- name: CreateAuthor :one
INSERT INTO authors (name)
VALUES (?)
RETURNING *;

-- name: CreateBookSubjectRelation :exec
INSERT INTO books_subjects (book_id, subject_id)
VALUES (?, ?);

-- name: CreateBookAuthorRelation :exec
INSERT INTO books_authors (book_id, author_id)
VALUES (?, ?);