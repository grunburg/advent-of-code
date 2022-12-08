package main

import (
	"bufio"
	"os"
)

func main() {
	one()
	two()
}

func scanner(input string) *bufio.Scanner {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file)
}
