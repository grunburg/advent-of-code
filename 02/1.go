package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("02/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// 1 Rock (A, X)
	// 2 Paper (B, Y)
	// 3 Scissors (C, Z)

	// 0 Lost
	// 6 Win
	// 3 Draw

	results := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,

		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,

		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		total = total + results[line]
	}

	fmt.Println(total)
}
