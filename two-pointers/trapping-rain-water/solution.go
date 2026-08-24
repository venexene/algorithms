package main

import "fmt"

func main() {
	height := []int{3, 0, 0, 5}
	result := trap(height)
	fmt.Println(result)
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left := 0
	right := len(height)-1

	leftMax := height[left]
	rightMax := height[right]

	sum := 0
	for left < right {
		if leftMax < rightMax {
			leftMax = max(leftMax, height[left])
			sum += leftMax - height[left]
			left++
		} else {
			rightMax = max(rightMax, height[right])
			sum += rightMax - height[right]
			right--
		}
	}

	return sum
}