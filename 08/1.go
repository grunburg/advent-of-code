package main

import (
	"fmt"
)

func one() {
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

	visibleTrees := 0
	for i := 0; i < len(treeHeights); i++ {
		for j := 0; j < len(treeHeights[i]); j++ {
			if isTreeVisible(treeHeights, i, j) {
				visibleTrees++
			}
		}
	}

	fmt.Println("Visible trees:", visibleTrees)
}

func isTreeVisible(treeHeights [][]int, row int, col int) bool {
	// Check if the tree is on the edge
	if row == 0 || col == 0 || row == len(treeHeights[row])-1 || col == len(treeHeights[col])-1 {
		return true
	}

	// Check if the tree is visible from the top edge
	temp := 0
	for i := row - 1; i >= 0; i-- {
		if treeHeights[i][col] > temp {
			temp = treeHeights[i][col]
		}
	}
	if treeHeights[row][col] > temp {
		return true
	}

	// Check if the tree is visible from the right edge
	temp = 0
	for i := col + 1; i < len(treeHeights[col]); i++ {
		if treeHeights[row][i] > temp {
			temp = treeHeights[row][i]
		}
	}
	if treeHeights[row][col] > temp {
		return true
	}

	// Check if the tree is visible from the bottom edge
	temp = 0
	for i := row + 1; i < len(treeHeights[row]); i++ {
		if treeHeights[i][col] > temp {
			temp = treeHeights[i][col]
		}
	}
	if treeHeights[row][col] > temp {
		return true
	}

	// Check if the tree is visible from the left edge
	//fmt.Println("l", row, col)
	temp = 0
	for i := col - 1; i >= 0; i-- {
		if treeHeights[row][i] > temp {
			temp = treeHeights[row][i]
		}
	}
	if treeHeights[row][col] > temp {
		return true
	}

	return false
}
