package main

import "fmt"

func main() {
	x := 7
	result := mySqrt(x)
	fmt.Println(result)
}

func mySqrt(x int) int {
	low := 0
	high := x

	for low <= high {
		mid := (high + low) / 2
		sqr := mid * mid

		if sqr > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low - 1
}
