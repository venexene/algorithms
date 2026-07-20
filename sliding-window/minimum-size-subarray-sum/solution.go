package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 11
	result := minSubArrayLen(target, nums)
	fmt.Println(result)
}

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	res := len(nums) + 1
	sum := 0

	for right < len(nums) {
		sum += nums[right]
		if sum >= target {
			res = min(res, right - left + 1)
			for sum >= target {
				res = min(res, right - left + 1)
				sum -= nums[left]
				left++
			}
		}
		right++
	}

	if res == len(nums) + 1 {
		return 0
	}

	return res
}