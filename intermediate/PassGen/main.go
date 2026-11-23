package main

import (
	"fmt"

	password "github.com/eduardoteranisi/study-go/intermediate/PassGen/src"
)

func main() {
	fmt.Println("Welcome to password generator CLI")

	p, err := password.New()
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
