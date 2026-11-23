package password

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eduardoteranisi/study-go/utils"
)

type Password struct {
	length    int
	upperCase bool
	lowerCase bool
	numbers   bool
	symbols   bool
}

func New() (*Password, error) {
	var p Password
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Type the wanted length for your password: ")
	l, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	lgth, err := strconv.Atoi(strings.TrimSpace(l))
	if err != nil {
		return nil, err
	}
	p.length = lgth

	fmt.Println("\nInclude upper case letters?")
	p.upperCase = utils.BoolOpt()

	fmt.Println("\nInclude lower case letters?")
	p.lowerCase = utils.BoolOpt()

	fmt.Println("\nInclude numbers?")
	p.numbers = utils.BoolOpt()

	fmt.Println("\nInclude symbols?")
	p.symbols = utils.BoolOpt()

	return &p, nil
}
