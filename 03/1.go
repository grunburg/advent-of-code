package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSolutionOne() {
	file, err := os.Open("03/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	// Solution 1
	var letters = make(map[int][]string)

	for index, line := range input {
		first, second := line[:len(line)/2], line[len(line)-len(line)/2:]

		for _, a := range first {
			for _, b := range second {
				if a == b {
					if includes(letters[index], string(a)) {
						continue
					}
					letters[index] = append(letters[index], string(a))
				}
			}
		}
	}

	var alphabet = make(map[string]int)
	priority := 1

	for i := 'a'; i <= 'z'; i++ {
		alphabet[string(i)] = priority
		priority++
	}

	for i := 'A'; i <= 'Z'; i++ {
		alphabet[string(i)] = priority
		priority++
	}

	total := 0

	for _, array := range letters {
		for _, letter := range array {
			total = total + alphabet[letter]
		}
	}

	fmt.Println(total)
}
