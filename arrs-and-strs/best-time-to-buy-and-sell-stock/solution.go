package main

import (
	"fmt"
)

func main() {
	prices := []int{7,6,4,3,1}
	result := maxProfitDynamic(prices)
	fmt.Println(result)
}

func maxProfitBrutforce(prices []int) int {
	max := 0
    for i := 0; i < len(prices) - 1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] - prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

func maxProfitDynamic(prices []int) int {
	min := prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] - min > profit {
			profit = prices[i] - min
		}
		
		if prices[i] < min {
			min = prices[i]
		}
	}

	return profit
}