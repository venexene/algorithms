package main

import "fmt"

func main() {
	nums := []int{4,10,4,3,8,9}
	result := lengthOfLIS(nums)
	fmt.Println(result)
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		res = max(dp[i], res)
	}

	return res
}