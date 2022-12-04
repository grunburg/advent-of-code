package main

import "fmt"

func one() {
	var input []string

	scanner := scanner("04/input.txt")
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	counter := 0

	split := split(input)
	for _, sections := range split {
		a := sections[0]
		b := sections[1]

		if a.last <= b.last && a.first >= b.first || a.last >= b.last && a.first <= b.first {
			counter++
		}
	}

	fmt.Println(counter)
}
