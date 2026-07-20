package main

import "fmt"

func main() {
	height := []int{8, 7, 2, 1}
	result := maxArea(height)
	fmt.Println(result)
}

func maxArea(height []int) int {
	left := 0
    right := len(height) - 1
	res := 0

	for left < right {
		res = max(res, (right-left)*min(height[left], height[right]))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return res
}