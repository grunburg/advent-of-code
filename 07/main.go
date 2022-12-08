package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var input []string
var sizes = make(map[string]int)
var current = 0

func main() {
	scanner := scanner("07/input.txt")
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	recursive("/")

	// Solution 1
	fmt.Println(getTotalSize())

	// Solution 2
	fmt.Println(getSmallestSize())
}

func scanner(input string) *bufio.Scanner {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file)
}

func recursive(path string) {
	size := 0
	for current < len(input) {
		line := input[current]
		if strings.HasPrefix(line, "dir") {
			current++
			continue
		}
		if strings.HasPrefix(line, "$ cd ") {
			cd := strings.TrimPrefix(line, "$ cd ")
			if cd == ".." {
				break
			}
			current += 2
			recursive(path + "/" + cd)
			size += sizes[path+"/"+cd]
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

func getSmallestSize() int {
	free := 70000000 - sizes["/"]
	need := 30000000 - free

	var temp []int
	for _, size := range sizes {
		temp = append(temp, size)
	}

	result := 0

	sort.Ints(temp)
	for _, size := range temp {
		if size >= need {
			result = size
			break
		}
	}

	return result
}

func getTotalSize() int {
	sum := 0
	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}
