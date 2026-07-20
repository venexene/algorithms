package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "a"
	result := checkInclusion(s1, s2)
	fmt.Println(result)
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	arr := [26]int{}
	for _, r := range s1 {
		arr[r-'a']++
	}

	need := 0
	for _, r := range arr {
		if r != 0 {
			need++
		}
	}

	count := 0

	for i := 0; i < len(s1); i++ {
		arr[rune(s2[i])-'a']--
		if arr[rune(s2[i])-'a'] == 0 {
			count++
		}
	}

	if count == need {
		return true
	}

	left := 1
	right := len(s1)
	for right < len(s2) {
		prevL := arr[rune(s2[left-1])-'a']
		arr[rune(s2[left-1])-'a']++
		if prevL == 0 && arr[rune(s2[left-1])-'a'] != 0 {
			count--
		}

		prevR := arr[rune(s2[right])-'a']
		arr[rune(s2[right])-'a']--
		if prevR != 0 && arr[rune(s2[right])-'a'] == 0 {
			count++
		}

		if count == need {
			return true
		}
		right++
		left++
	}

	return false
}

func checkInclusionMap(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	mp := map[rune]int{}
	for _, r := range s1 {
		mp[r]++
	}

	for i := 0; i < len(s1); i++ {
		if _, ok := mp[rune(s2[i])]; ok {
			mp[rune(s2[i])]--
		}
	}

	allZero := true
	for _, v := range mp {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return true
	}

	left := 1
	right := len(s1)
	for right < len(s2) {
		if _, ok := mp[rune(s2[left-1])]; ok {
			mp[rune(s2[left-1])]++
		}
		if _, ok := mp[rune(s2[right])]; ok {
			mp[rune(s2[right])]--
		}

		allZero = true
		for _, v := range mp {
			if v != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			return true
		}

		right++
		left++
	}

	return false
}
