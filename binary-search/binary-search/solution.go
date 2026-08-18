package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	result := search(nums, target)
	fmt.Println(result)
}

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (high + low) / 2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
