package main

import "fmt"

func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}

func partitionLabels(s string) []int {
	runes := []rune(s)
	last := map[rune]int{}

	for i, v := range(runes) {
		last[v] = i
	}

	res := []int{}
	start := 0
	end := 0
	for i, v := range(runes) {
		end = max(end, last[v])
		if end == i {
			res = append(res, end - start + 1)
			start = i + 1
		}
	}

	return res
}