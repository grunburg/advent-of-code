package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSolutionTwo() {
	file, err := os.Open("03/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var rucksacks = make(map[int][]string)
	counter := 0
	group := 1

	for scanner.Scan() {
		line := scanner.Text()
		rucksacks[group] = append(rucksacks[group], line)

		counter++

		if counter%3 == 0 {
			group++
		}
	}

	var same []string

	for _, grouped := range rucksacks {
		characters := getSameCharacterArray(grouped[0], grouped[1], grouped[2])
		for _, char := range characters {
			same = append(same, char)
		}
	}

	total := getTotal(same)
	fmt.Println(total)
}

// array //array
func getSameCharacterArray(a string, b string, c string) []string {
	var iteration1 []string

	for _, first := range a {
		for _, second := range b {
			if first == second {
				if includes(iteration1, string(first)) {
					continue
				}
				iteration1 = append(iteration1, string(first))
			}
		}
	}

	var iteration2 []string

	for _, first := range c {
		for _, second := range iteration1 {
			if string(first) == second {
				if includes(iteration2, string(first)) {
					continue
				}
				iteration2 = append(iteration2, string(first))
			}
		}
	}

	return iteration2
}

func getTotal(a []string) int {
	alphabet := getAlphabet()
	total := 0

	for _, array := range a {
		for _, letter := range array {
			total = total + alphabet[string(letter)]
		}
	}

	return total
}

func getAlphabet() map[string]int {
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

	return alphabet
}
