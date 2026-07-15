package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "anagram"
	t := "nagaram"
	result := isAnagramHash(s, t)
	fmt.Println(result)
}

func isAnagramArr(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	symbols := [26]int{}

	for i := range s {
		symbols[s[i] - 'a']++
		symbols[t[i] - 'a']--
	}

	for _, c := range symbols {
		if c != 0 {
			return false
		}
	}

	return true
}

func isAnagramHash(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[byte]int{}
	tMap := map[byte]int{}

	for i := range s {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	for k := range sMap {
		if sMap[k] != tMap[k] {
			return false
		}
	}

	return true
}

func isAnagramSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := sortString(s)
	tr := sortString(t)

	for i, symb := range sr {
		if tr[i] != symb {
			return false
		}
	}
	return true
}

func sortString(s string) []rune {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return r
}