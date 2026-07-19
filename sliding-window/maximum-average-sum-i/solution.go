package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	res := findMaxAverage(nums, k)
	fmt.Println(res)
}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	maxSum := 0
    for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum = sum

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i - k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}