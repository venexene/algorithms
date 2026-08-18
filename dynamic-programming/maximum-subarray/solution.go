package main

import "fmt"

func main() {
	nums := []int{5, 4, -1, 7, 8}
	result := maxSubArray(nums)
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i]+nums[i-1])
		sum = max(nums[i], sum)
	}

	return sum
}
