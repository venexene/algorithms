package main

func main() {

}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	queue := [][]int{}
	orangesNum := 0
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				orangesNum++
			}
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	if orangesNum == 0 {
		return 0
	}

	minute := -1
	for len(queue) > 0 {
		orangesInMinute := len(queue)
		for l := 0; l < orangesInMinute; l++ {
			rotten := queue[0]
			queue = queue[1:]

			i := rotten[0]
			j := rotten[1]
			if i+1 < n && grid[i+1][j] == 1 {
				grid[i+1][j] = 2
				orangesNum--
				queue = append(queue, []int{i + 1, j})
			}
			if i-1 >= 0 && grid[i-1][j] == 1 {
				grid[i-1][j] = 2
				orangesNum--
				queue = append(queue, []int{i - 1, j})
			}
			if j+1 < m && grid[i][j+1] == 1 {
				grid[i][j+1] = 2
				orangesNum--
				queue = append(queue, []int{i, j + 1})
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				grid[i][j-1] = 2
				orangesNum--
				queue = append(queue, []int{i, j - 1})
			}
		}
		minute++
	}

	if orangesNum == 0 {
		return minute
	}
	return -1
}
