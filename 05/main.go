package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	one()
	two()
}

type Instruction struct {
	amount int
	from   int
	to     int
}

func structure() map[int][]string {
	return map[int][]string{
		1: {"Z", "J", "G"},
		2: {"Q", "L", "R", "P", "W", "F", "V", "C"},
		3: {"F", "P", "M", "C", "L", "G", "R"},
		4: {"L", "F", "B", "W", "P", "H", "M"},
		5: {"G", "C", "F", "S", "V", "Q"},
		6: {"W", "H", "J", "Z", "M", "Q", "T", "L"},
		7: {"H", "F", "S", "B", "V"},
		8: {"F", "J", "Z", "S"},
		9: {"M", "C", "D", "P", "F", "H", "B", "T"},
	}
}

func scanner(input string) *bufio.Scanner {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file)
}

func instructions(raw []string) []Instruction {
	var instructions []Instruction

	for _, line := range raw {
		regex := regexp.MustCompile(`\d+`)
		numbers := regex.FindAllString(line, -1)

		instructions = append(instructions, Instruction{amount: convert(numbers[0]), from: convert(numbers[1]), to: convert(numbers[2])})
	}

	return instructions
}

func convert(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return i
}

func build(raw map[int][]string) map[int]*list.List {
	stacks := make(map[int]*list.List)

	for index, values := range raw {
		a := list.New()
		for _, value := range values {
			a.PushFront(value)
		}
		stacks[index] = a
	}

	return stacks
}

func stack(stacks map[int]*list.List) {
	for index, stack := range stacks {
		fmt.Printf("%d ", index)
		for e := stack.Back(); e != nil; e = e.Prev() {
			fmt.Printf("%s ", e.Value)
		}
		fmt.Println()
	}
}
