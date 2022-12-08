package main

import (
	"log"
)

func one() {
	var input []string

	scanner := scanner("07/input.txt")
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	recursive("/", input)

	sum := 0
	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}

	log.Println(sum)
}
