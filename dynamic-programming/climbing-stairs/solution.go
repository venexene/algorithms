package main

import "fmt"

func main() {
	n := 2
	result := climbStairs(n)
	fmt.Println(result)
}

func climbStairs(n int) int {
	s1 := 1
	s2 := 1
    for i := 0; i < n; i++ {
		s1, s2 = s2, s1+s2
	}
	return s1
}
