package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
    top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			for j := right; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
			bottom--
		}
	
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
		
	}
	
	return res
}