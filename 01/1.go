package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func top(m map[int]int) []int {
	var values []int

	for _, value := range m {
		values = append(values, value)
	}

	sort.Ints(values)
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	return values
}

func main() {
	file, _ := os.Open("01/elfs.txt")

	scanner := bufio.NewScanner(file)

	elfs := make(map[int][]int)
	elf := 1

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			elf++
			continue
		}

		i, _ := strconv.Atoi(line)
		elfs[elf] = append(elfs[elf], i)
	}

	totals := make(map[int]int)

	for elf, calories := range elfs {
		total := 0
		for _, calorie := range calories {
			total += calorie
		}
		totals[elf] = total
	}

	// Return sorted array of elfs with the most calories
	descending := top(totals)

	fmt.Printf("Top 3: %d\n", descending[0]+descending[1]+descending[2])
	fmt.Printf("Top 1: %d\n", descending[0])
}
