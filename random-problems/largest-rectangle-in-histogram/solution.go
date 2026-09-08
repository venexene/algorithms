package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	res := largestRectangleArea(heights)
	fmt.Println(res)
}

func largestRectangleArea(heights []int) int {
	left := make([]int, len(heights))
	stack := []int{}
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	right := make([]int, len(heights))
	stack = []int{}
	for i := len(heights)-1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			right[i] = len(heights)
		} else {
			right[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}
	
	res := 0
	for i := 0; i < len(heights); i++ {
		area := heights[i] * (right[i] - left[i] - 1)
		res = max(area, res)
	}

	return res
}