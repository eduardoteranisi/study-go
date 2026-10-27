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

-- name: GetBookByTitle :one
SELECT *
FROM books b
WHERE b.title = (?);

-- name: GetBookByAuthor :one
SELECT b.*, a.name
FROM books b
LEFT JOIN books_authors ba ON b.id = ba.book_id
LEFT JOIN authors a ON ba.author_id = a.id
WHERE a.name = (?);

-- name: GetBookByIsbn :one
SELECT *
FROM books b
WHERE b.isbn = (?);