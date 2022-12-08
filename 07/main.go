package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sizes = make(map[string]int)
var current = 0

func main() {
	one()
}

func scanner(input string) *bufio.Scanner {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file)
}

func recursive(path string, commands []string) {
	size := 0
	for current < len(commands) {
		line := commands[current]
		if commands[current] == "$ cd .." {
			break
		}

		if strings.HasPrefix(line, "dir") {
			current++
			continue
		}
		if strings.HasPrefix(line, "$ cd ") {
			current += 2
			recursive(path+"/"+strings.TrimPrefix(line, "$ cd "), commands)
			size += sizes[path+"/"+strings.TrimPrefix(line, "$ cd ")]
		} else {
			size += convert(strings.Split(line, " ")[0])
		}
		current++
	}
	sizes[path] = size
}

func convert(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return i
}
