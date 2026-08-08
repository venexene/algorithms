package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','1','0'},
	}
	result := numIslands(grid)
	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	res := 0
    for r, row := range grid {
		for c, elem := range row {
			if elem == '1' {
				res++
				flood(grid, r, c)
			}
		}
	}
	return res
}

func flood(grid [][]byte, r, c int) {
	w := len(grid)
	h := len(grid[0])

	grid[r][c] = '0'
	if r + 1 < w && grid[r + 1][c] == '1' {
		flood(grid, r + 1, c)
	}
	if r - 1 >= 0 && grid[r - 1][c] =='1' {
		flood(grid, r - 1, c)
	}
	if c + 1 < h && grid[r][c + 1] == '1' {
		flood(grid, r, c + 1)
	}
	if c - 1 >= 0 && grid[r][c - 1] == '1' {
		flood(grid, r, c - 1)
	}
}