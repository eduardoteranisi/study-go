package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	// ctx := context.Background()

	fmt.Println("Welcome to MyNotesAppCLI")
	fmt.Println("What do you want to do?")
	for {
		option, err := userMenu()
		if err != nil {
			panic(err)
		}

		switch Options(option) {
		case CREATE:
			fmt.Println("Create option selected")
		case EDIT:
			fmt.Println("Edit option selected")
		case GET:
			fmt.Println("Get option selected")
		case DELETE:
			fmt.Println("Delete option selected")
		case EXIT:
			fmt.Println("Goodbye...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}

	// client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Disconnect(ctx)
}

func userMenu() (int, error) {
	fmt.Println()
	fmt.Println("========== MENU ==========")
	fmt.Println("1 - Create a note")
	fmt.Println("2 - Edit a note")
	fmt.Println("3 - Get a note")
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
