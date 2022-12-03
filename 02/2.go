package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSolutionTwo() {
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

	// X Lose
	// Y Draw
	// Z Win

	results := map[string]int{
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,

		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,

		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		total = total + results[line]
	}

	fmt.Println(total)
}
