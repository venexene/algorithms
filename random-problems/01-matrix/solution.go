package main

import "fmt"

func main() {
	mat := [][]int{
		{1, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
	}
	result := updateMatrix(mat)
	fmt.Println(result)
}

func updateMatrix(mat [][]int) [][]int {
    res := make([][]int, len(mat))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, 0, len(mat[0]))
		for j := 0; j < len(mat[0]); j++ {
			res[i] = append(res[i], -1)
		} 	
	}

	queue := [][]int{}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
				queue = append(queue, []int{i, j})
			}
		}
	}
	bfs(mat, res, queue) 

	return res
}

func bfs(mat, res, queue [][]int) {
	n := len(mat)
	m := len(mat[0])
	lenPath := 0
	for len(queue) > 0 {
		ln := len(queue)
		lenPath++
		for i := 0; i < ln; i++ {
			pos := queue[0]
			queue = queue[1:]

			i := pos[0]
			j := pos[1]
			if i+1 < n && res[i+1][j] == -1 && mat[i+1][j] != 0 {
				res[i+1][j] = lenPath
				queue = append(queue, []int{i+1, j})
			}
			if i-1 >= 0 && res[i-1][j] == -1 && mat[i-1][j] != 0 {
				res[i-1][j] = lenPath
				queue = append(queue, []int{i-1, j})
			}
			if j+1 < m && res[i][j+1] == -1 && mat[i][j+1] != 0 {
				res[i][j+1] = lenPath
				queue = append(queue, []int{i, j+1})
			}
			if j-1 >= 0 && res[i][j-1] == -1 && mat[i][j-1] != 0 {
				res[i][j-1] = lenPath
				queue = append(queue, []int{i, j-1})
			}
		}
	}
}
