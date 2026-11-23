package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BoolOpt() bool {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1) Yes")
		fmt.Println("2) No")

		in, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		switch strings.TrimSpace(in) {
		case "1":
			return true
		case "2":
			return false
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
