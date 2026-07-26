package main

import "fmt"

func main() {
	piles := []int{30,11,23,4,20}
	h := 5
	result := minEatingSpeed(piles, h)
	fmt.Println(result)
}

func minEatingSpeed(piles []int, h int) int {
	low := 1
	high := 0
	for _, p := range piles {
		high = max(high, p) 
	}

	for low < high {
		mid := (low + high) / 2

		if countTime(piles, mid) <= h {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func countTime(piles []int, speed int) int {
	t := 0
	for _, p := range piles {
		t += (p + speed - 1) / speed
	}
	return t
}