package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eduardoteranisi/study-go/beginner/NotesApp/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const mongoURI = "mongodb://localhost:27017"

type Options int

const (
	EXIT Options = iota
	CREATE
	EDIT
	GET
	DELETE
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := client.Database("notes_db")
	err = db.CreateCollection(ctx, "notes")
	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to MyNotesAppCLI")
	fmt.Println("What do you want to do?")
	for {
		option, err := userMenu()
		if err != nil {
			panic(err)
		}

		switch Options(option) {
		case CREATE:
			fmt.Println("\nCreate option selected")
			services.CreateNote(ctx, db)
		case EDIT:
			fmt.Println("\nEdit option selected")
			services.EditNote(ctx, db)
		case GET:
			fmt.Println("\nGet option selected")
			services.GetNote(ctx, db)
		case DELETE:
			fmt.Println("\nDelete option selected")
			services.DeleteNote(ctx, db)
		case EXIT:
			fmt.Println("Goodbye...")
			return
		default:
			fmt.Println("\nInvalid option")
		}
	}

}

func userMenu() (int, error) {
	fmt.Println()
	fmt.Println("========== MENU ==========")
	fmt.Println("1 - Create note")
	fmt.Println("2 - Edit note")
	fmt.Println("3 - Get notes")
	fmt.Println("4 - Delete note")
	fmt.Println("0 - Exit")
	fmt.Println("==========================")
	fmt.Print("Choose an option: ")

	r := bufio.NewReader(os.Stdin)
	in, err := r.ReadString('\n')
	if err != nil {
		return -1, err
	}
	in = strings.TrimSpace(in)

	opt, err := strconv.Atoi(in)
	if err != nil {
		return -1, err
	}

	return opt, nil
}
