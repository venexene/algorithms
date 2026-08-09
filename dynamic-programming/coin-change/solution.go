package main

import "fmt"

func main() {
	coins := []int{2}
	amount := 3
	result := coinChange(coins, amount)
	fmt.Println(result)
}

func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				if dp[i] == -1 || 1 + dp[i-coin] < dp[i] {
					dp[i] = 1 + dp[i - coin]
				}
			}
		}
	}

	return dp[amount]
}