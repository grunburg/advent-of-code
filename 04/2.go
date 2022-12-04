package main

import "fmt"

func two() {
	var input []string

	scanner := scanner("04/input.txt")
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	counter := 0

	split := split(input)
	for _, sections := range split {
		var a []int
		for i := sections[0].first; i <= sections[0].last; i++ {
			a = append(a, i)
		}

		var b []int
		for i := sections[1].first; i <= sections[1].last; i++ {
			b = append(b, i)
		}

		for _, value := range a {
			if contains(b, value) {
				counter++
				break
			}
		}
	}

	fmt.Println(counter)
}
