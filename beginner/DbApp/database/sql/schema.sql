CREATE TABLE IF NOT EXISTS books (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    isbn TEXT NOT NULL,
    title TEXT NOT NULL,
    subtitle TEXT,
    publisher TEXT,
    year INT,
    format TEXT,
    page_count INT
);

CREATE TABLE IF NOT EXISTS subjects (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS authors (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS books_subjects (
    book_id   INTEGER NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    subject_id INTEGER NOT NULL REFERENCES subjects(id),
    PRIMARY KEY (book_id, subject_id)
);

CREATE TABLE IF NOT EXISTS books_authors (
    book_id   INTEGER NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    author_id INTEGER NOT NULL REFERENCES authors(id),
    PRIMARY KEY (book_id, author_id)
);