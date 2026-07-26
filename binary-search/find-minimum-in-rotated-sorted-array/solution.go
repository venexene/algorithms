package main

import "fmt"

func main() {
	nums := []int{5, 1, 2, 3, 4}
	result := findMin(nums)
	fmt.Println(result)
}

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		mid := (low + high) / 2

		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		}
	}

	return nums[low]
}
