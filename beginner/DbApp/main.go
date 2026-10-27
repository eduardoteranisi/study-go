package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/eduardoteranisi/study-go/beginner/DbApp/database"
	"github.com/eduardoteranisi/study-go/beginner/DbApp/database/sqlc"
	"github.com/eduardoteranisi/study-go/beginner/DbApp/services"
)

const baseURL = "https://brasilapi.com.br/api/isbn/v1/"
const paramsURL = "?providers="
const providers = "open-library"
const dbPath = "file:./database/app.db?_foreign_keys=on"

type Book struct {
	Isbn      string   `json:"isbn"`
	Title     string   `json:"title"`
	Subtitle  string   `json:"subtitle"`
	Authors   []string `json:"authors"`
	Publisher string   `json:"publisher"`
	Year      int      `json:"year"`
	Format    string   `json:"format"`
	PageCount int      `json:"page_count"`
}

func main() {
	ctx := context.Background()

	fmt.Print("Digite o ISBN do livro: ")

	osIn := bufio.NewReader(os.Stdin)

	isbn, err := osIn.ReadString('\n')
	if err != nil {
		panic(err)
	}
	isbn = strings.TrimSpace(isbn)

	endpoint := baseURL + url.PathEscape(isbn) + paramsURL + providers

	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bookInfo, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var bookData Book
	err = json.Unmarshal(bookInfo, &bookData)
	if err != nil {
		panic(err)
	}

	db, err := database.Open(dbPath)
	if err != nil {
		panic(err)
	}

	defer database.Close(db)

	err = database.Migrate(db)
	if err != nil {
		panic(err)
	}

	createdBook, err := database.CreateBook(ctx, db, sqlc.CreateBookParams{
		Isbn:      bookData.Isbn,
		Title:     bookData.Title,
		Subtitle:  nullStringHelper(bookData.Subtitle),
		Publisher: nullStringHelper(bookData.Publisher),
		Year:      nullInt64Helper(bookData.Year),
		Format:    nullStringHelper(bookData.Format),
		PageCount: nullInt64Helper(bookData.PageCount),
	})
	if err != nil {
		panic(err)
	}

	services.ReadSubjectFromUserService(ctx, db, createdBook.ID)
	services.CreateAuthorsService(ctx, db, bookData.Authors, createdBook.ID)

	fmt.Printf("Book added to db: %s", createdBook.Title)
}

func nullStringHelper(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func nullInt64Helper(n int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(n), Valid: n != 0}
}
