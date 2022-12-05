package main

func two() {
	var input []string

	scanner := scanner("05/input.txt")
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	stacks := build(structure())
	instructions := instructions(input)
	for _, instruction := range instructions {
		var removes []any

		// Remove
		for i := 1; i <= instruction.amount; i++ {
			e := stacks[instruction.from].Front()
			removes = append(removes, e.Value)
			stacks[instruction.from].Remove(e)
		}

		removes = reverse(removes)

		// Add
		for _, element := range removes {
			stacks[instruction.to].PushFront(element)
		}
	}

	// Solution
	out(stacks)
}

func reverse(array []any) []any {
	last := len(array) - 1
	for i := 0; i < len(array)/2; i++ {
		array[i], array[last-i] = array[last-i], array[i]
	}

	return array
}
