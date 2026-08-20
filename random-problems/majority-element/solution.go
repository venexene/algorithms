package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement(nums)
	fmt.Println(res)
}

func majorityElement(nums []int) int {
    major := 0
	c := 0
	for _, n := range nums {
		if c == 0 {
			major = n
			c = 1
		} else if n == major {
			c++
		} else {
			c--
		}
	}
	return major
}
