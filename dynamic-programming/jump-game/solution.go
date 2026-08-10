package main

import "fmt"

func main() {

}

func canJumpUnoptimized(nums []int) bool {
    dp := make([]bool, len(nums))

	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if !dp[i] {
			continue
		}
		p := nums[i]
		for j := p; j > 0; j-- {
			fmt.Println(i + j)
			if i + j == len(nums) - 1 {
				return true
			}
			if i + j < len(nums) - 1 {
				dp[i+j] = true
			}
		}
	}


	fmt.Println(dp)
	return dp[len(nums) - 1]
}

func canJump(nums []int) bool {
	mx := 0

	for i := 0; i < len(nums); i++ {
		if mx >= len(nums) - 1 {
			return true
		}
		if mx >= i {
			mx = max(mx, i + nums[i])
		}
	}

	return mx >= len(nums) - 1
}