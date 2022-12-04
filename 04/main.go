package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

type Section struct {
	first int
	last  int
}

func scanner(input string) *bufio.Scanner {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file)
}

func split(raw []string) map[int][]Section {
	var result = make(map[int][]Section)

	for i, line := range raw {
		split := strings.Split(line, ",")

		for _, a := range split {
			b := strings.Split(a, "-")
			section := Section{first: convert(b[0]), last: convert(b[1])}
			result[i] = append(result[i], section)
		}
	}

	return result
}

func convert(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return i
}

func contains(array []int, value int) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}
