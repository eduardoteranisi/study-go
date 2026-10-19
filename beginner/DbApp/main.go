package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/eduardoteranisi/study-go/beginner/DbApp/database"
)

const baseURL = "https://brasilapi.com.br/api/isbn/v1/"
const dbPath = "./database/app.db"

func main() {
	fmt.Print("Digite o ISBN do livro: ")

	osIn := bufio.NewReader(os.Stdin)

	isbn, err := osIn.ReadString('\n')
	if err != nil {
		panic(err)
	}
	isbn = strings.TrimSpace(isbn)

	endpoint := baseURL + url.PathEscape(isbn)

	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", body)

	db, err := database.Open(dbPath)
	if err != nil {
		panic(err)
	}

	err = database.Migrate(db)
	if err != nil {
		panic(err)
	}
}
