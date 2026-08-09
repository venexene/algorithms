package main

func main() {

}

func pacificAtlantic(heights [][]int) [][]int {
	if heights == nil {
		return nil
	}
	if len(heights) == 0 {
		return [][]int{}
	}

	m := len(heights)
	n := len(heights[0])

	pacific := [][]bool{}
	for i := 0; i < m; i++ {
		dfs(heights, pacific, m, n, i, 0)
	}
	for i := 0; i < n; i++ {
		dfs(heights, pacific, m, n, 0, i)
	}

	atlantic := [][]bool{}
	for i := m - 1; i >= 0; i-- {
		dfs(heights, atlantic, m, n, i, n - 1)
	}
	for i := n - 1; i >= 0; i-- {
		dfs(heights, atlantic, m, n, m - 1, i)
	}

	res := [][]int{}
	for k := range pacific {
		if atlantic[k] {
			res = append(res, []int{k[0], k[1]})
		}
	}

	return res
}

func dfs(heights [][]int, visited [][]bool, m, n, r, c int) {
	visited[r][c] = true
	if r + 1 < m && !visited[r+1][c] && heights[r+1][c] >= heights[r][c] {
		dfs(heights, visited, m, n, r + 1, c)
	}
	if r - 1  >= 0 && !visited[r-1][c] && heights[r-1][c] >= heights[r][c] {
		dfs(heights, visited, m, n, r - 1, c)
	}
	if c + 1 < n && !visited[r][c+1] && heights[r][c+1] >= heights[r][c] {
		dfs(heights, visited, m, n, r, c + 1)
	}
	if c - 1  >= 0 && !visited[r][c-1] && heights[r][c-1] >= heights[r][c] {
		dfs(heights, visited, m, n, r, c - 1)
	}
}