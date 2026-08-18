package main

import "fmt"

func main() {
	m := 3
	n := 7
	result := uniquePaths(m, n)
	fmt.Println(result)
}

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for i, r := range grid {
		for j := range r {
			if i == 0 || j == 0 {
				grid[i][j] = 1
				continue
			}
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}
