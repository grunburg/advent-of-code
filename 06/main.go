package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	solve(input, 4)
	solve(input, 14)
}

func solve(input string, receiver int) {
	for i := 0; i < len(input)-receiver; i++ {
		part := input[i : i+receiver]
		if isAllDifferent(part) {
			fmt.Println(part, i+receiver)
			return
		}
	}
}

func isAllDifferent(part string) bool {
	seen := make(map[rune]bool)

	for _, char := range part {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}
