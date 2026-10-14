package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader_in := bufio.NewReader(os.Stdin)
	var out int

	str_in, err := reader_in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	rune_conv_in := []rune(strings.TrimSpace(str_in))

	size_bits := len(rune_conv_in)

	digits := make([]int, size_bits)

	for i := range size_bits {
		str_digit := string(rune_conv_in[i])

		digits[i], err = strconv.Atoi(str_digit)
		if err != nil {
			fmt.Println(err)
		}
	}

	slices.Reverse(digits)

	for i, digit := range digits {
		out += digit * int(math.Pow(2, float64(i)))
	}

	fmt.Println(out)
}
