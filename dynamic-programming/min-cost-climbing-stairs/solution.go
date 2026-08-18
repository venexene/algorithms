package main

import "fmt"

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	result := minCostClimbingStairs(cost)
	fmt.Println(result)
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	} else if len(cost) == 1 {
		return cost[0]
	} else if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}

	return min(cost[len(cost)-1], cost[len(cost)-2])
}
