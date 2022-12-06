package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	getResult(input, 4)
	getResult(input, 14)
}

func getResult(input string, receiver int) {
	for i := 0; i < len(input)-receiver; i++ {
		part := input[i : i+receiver]
		if isPartDifferent(part) {
			out(part, i+receiver)
			return
		}
	}
}

func isPartDifferent(part string) bool {
	seen := make(map[rune]bool)

	for _, char := range part {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func out(part string, count int) {
	fmt.Printf("Part: %s, Count: %d\n", part, count)
}
