package main

import (
	"fmt"
	"sort"
)

func two() {
	var rows []string

	scanner := scanner("08/input.txt")
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, line)
	}

	cols := len(rows[0])
	treeHeights := make([][]int, len(rows))

	for i := 0; i < len(rows); i++ {
		for j := 0; j < cols; j++ {
			treeHeights[i] = append(treeHeights[i], int(rows[i][j]-'0'))
		}
	}

	var scores []int
	for i := 0; i < len(treeHeights); i++ {
		for j := 0; j < len(treeHeights[i]); j++ {
			scores = append(scores, getScenicScore(treeHeights, i, j))
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	fmt.Println("Highest scenic score:", scores[0])
}

func getScenicScore(treeHeights [][]int, row int, col int) int {
	t := 0
	for i := row - 1; i >= 0; i-- {
		t++
		if treeHeights[i][col] >= treeHeights[row][col] {
			break
		}
	}

	r := 0
	for i := col + 1; i < len(treeHeights[col]); i++ {
		r++
		if treeHeights[row][i] >= treeHeights[row][col] {
			break
		}
	}

	b := 0
	for i := row + 1; i < len(treeHeights[row]); i++ {
		b++
		if treeHeights[i][col] >= treeHeights[row][col] {
			break
		}
	}

	l := 0
	for i := col - 1; i >= 0; i-- {
		l++
		if treeHeights[row][i] >= treeHeights[row][col] {
			break
		}
	}

	return t * r * b * l
}
