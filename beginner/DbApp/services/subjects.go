package services

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/eduardoteranisi/study-go/beginner/DbApp/database"
)

func ReadSubjectFromUserService(ctx context.Context, db *sql.DB, book_id int64) error {
	fmt.Print("Digite o(s) gênero(s) do livro: ")

	osIn := bufio.NewReader(os.Stdin)

	userSubj, err := osIn.ReadString('\n')
	if err != nil {
		return err
	}
	userSubj = strings.TrimSpace(userSubj)

	subjects := strings.Split(userSubj, ",")

	err = database.CreateSubjects(ctx, db, subjects, book_id)
	if err != nil {
		return err
	}

	return nil
}
