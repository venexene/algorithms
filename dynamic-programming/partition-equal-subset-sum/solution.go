package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2, 2, 3,5}
	result := canPartition(nums)
	fmt.Println(result)
}

func canPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total % 2 == 1 {
		return false
	}
	target := total / 2

	dp := make([]bool, target+1)
	dp[0] = true

	for _, n := range nums {
		for s := target; s >= n; s-- {
			dp[s] = dp[s] || dp[s - n]
		}
	}

	return dp[target]
}
