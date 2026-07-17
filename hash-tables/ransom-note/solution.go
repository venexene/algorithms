package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "aab"
	result := canConstructArr(ransomNote, magazine)
	fmt.Println(result)
}

func canConstructArr(ransomNote string, magazine string) bool {
	arr := [26]int{}

	for _, r := range magazine {
		arr[r-'a']++
	}

	for _, r := range ransomNote {
		arr[r-'a']--
		if arr[r-'a'] < 0 {
			return false
		}
	}

	return true
}

func canConstructMap(ransomNote string, magazine string) bool {
	mp := map[rune]int{}

	for _, r := range magazine {
		mp[r]++
	}

	for _, r := range ransomNote {
		mp[r]--
		if mp[r] < 0 {
			return false
		}
	}

	return true
}