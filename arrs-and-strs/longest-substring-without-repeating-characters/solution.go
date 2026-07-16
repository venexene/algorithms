package main

import "fmt"

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
    mp := map[rune]int{}
	left := 0
	right := 0
	max := 0
	runes := []rune(s)
	for right < len(runes) {
		if val, ok := mp[runes[right]]; ok && val >= left {
			left = val + 1
		}

		if right - left + 1 > max {
			max = right - left + 1
		}

		mp[runes[right]] = right
		right++
	}
	return max
}