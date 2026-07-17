package main

import "fmt"

func main() {
	s := "aadadaad"
	result := firstUniqChar(s)
	fmt.Println(result)
}

func firstUniqChar(s string) int {
	mp := map[rune]int{}
	for _, r := range s {
		mp[r]++
	}

	for i, r := range s {
		if mp[r] == 1 {
			return i
		}
	}
	
	return -1
}