package main

import "fmt"

func main() {
	nums := []int{1}
	target := 0
	result := searchInsert(nums, target)
	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
    low := 0
	high := len(nums) - 1

	mid := -1
	for low <= high {
		mid = (high + low) / 2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}