package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"bdddddddddd","bbbbbbbbbbc"}
	result := groupAnagramsFreq(strs)
	fmt.Println(result)
}

func groupAnagramsSort(strs []string) [][]string {
	mp := map[string][]string{}

	for _, str := range strs {
		key := string(sortString(str))
		mp[key] = append(mp[key], str)
	}

	vals := make([][]string, 0, len(mp))
	for _, val := range mp {
		vals = append(vals, val)
	}
	
	return vals
}

func sortString(s string) []rune {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return r
}

func groupAnagramsFreq(strs []string) [][]string {
	mp := map[[26]int][]string{}

	for _, str := range strs {
		key := getFreq(str)
		mp[key] = append(mp[key], str)
	}

	vals := make([][]string, 0, len(mp))
	for _, val := range mp {
		vals = append(vals, val)
	}
	
	return vals
}

func getFreq(s string) [26]int {
	arr := [26]int{}
	runes := []rune(s)
	for _, r := range runes {
		arr[int(r - 'a')]++
	}

	return arr
}