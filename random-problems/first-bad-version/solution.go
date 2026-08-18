package main

import "fmt"

func main() {
	n := 10
	result := firstBadVersion(n)
	fmt.Println(result)
}

var bad = 4

func isBadVersion(version int) bool { 
	return version >= bad 
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right - left) / 2
		if !isBadVersion(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
