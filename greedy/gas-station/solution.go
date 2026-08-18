package main

import "fmt"

func main() {
	gas := []int{3, 1, 1}
	cost := []int{1, 2, 2}
	result := canCompleteCircuit(gas, cost)
	fmt.Println(result)
}

func canCompleteCircuit(gas []int, cost []int) int {
	tank := 0
	mx := 0
	sm := 0
	res := -1

	for i := len(gas) - 1; i >= 0; i-- {
		dif := gas[i] - cost[i]
		tank += dif
		if tank >= mx {
			mx = tank
			res = i
		}
		sm += dif
	}

	if sm < 0 {
		return -1
	}
	return res
}

func canCompleteCircuitForward(gas []int, cost []int) int {
	tank := 0
	sm := 0
	res := 0

	for i := 0; i < len(gas); i++ {
		dif := gas[i] - cost[i]
		tank += dif
		if tank < 0 {
			tank = 0
			res = i + 1
		}
		sm += dif
	}

	if sm < 0 {
		return -1
	}
	return res
}
