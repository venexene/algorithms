package main

import (
	"fmt"
)

func main() {
	s := "cbbd"
	result := longestPalindromeExpand(s)
	fmt.Println(result)
}

func longestPalindromeExpand(s string) string {
	runes := []rune(s)
	maxl := 0
	maxr := 0
	l := 0
	r := 0
	for i := 0; i < len(runes); i++ {
		l, r = expand(i, i, runes)
		if r - l + 1 > maxr - maxl + 1 {
			maxr = r
			maxl = l
		}

		if i + 1 >= len(runes) {
			continue
		}
		l, r = expand(i, i + 1, runes)
		if r - l + 1 > maxr - maxl + 1 {
			maxr = r
			maxl = l
		}
	}
	return string(runes[maxl:maxr+1])
}

func expand(l, r int, runes []rune) (int, int) {
	for l >= 0 && r < len(runes) && runes[l] == runes[r]  {
		l--;
		r++;
	}
	return l+1, r-1
}
