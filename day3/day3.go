package main

import (
	"bufio"
	"fmt"
	"os"
)

func countTrees(grid [][]bool, right int, down int) int {
	maxX := len(grid[0])
	x := 0
	treeCount := 0

	for r := 0; r < len(grid); r += down {
		if grid[r][x] {
			treeCount++
		}

		x = (x + right) % maxX
	}

	return treeCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var grid [][]bool

	for scanner.Scan() {
		line := scanner.Text()

		row := make([]bool, 31)
		for i := 0; i < len(line); i++ {
			row[i] = (line[i] == '#')
		}
		grid = append(grid, row)
	}

	treeCount := countTrees(grid, 3, 1)

	fmt.Printf("Part 1: %d\n", treeCount)

	treeProduct := 1
	treeProduct *= countTrees(grid, 1, 1)
	treeProduct *= countTrees(grid, 3, 1)
	treeProduct *= countTrees(grid, 5, 1)
	treeProduct *= countTrees(grid, 7, 1)
	treeProduct *= countTrees(grid, 1, 2)

	fmt.Printf("Part 2: %d\n", treeProduct)
}
