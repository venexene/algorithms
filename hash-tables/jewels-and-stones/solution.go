package main

import "fmt"

func main() {
	jewels := "aA"
	stones := "aAAbbb"
	result := numJewelsInStones(jewels, stones)
	fmt.Println(result)
}

func numJewelsInStones(jewels string, stones string) int {
	set := map[rune]struct{}{}
	count := 0

	for _, j := range jewels {
		set[j] = struct{}{}
	}

	for _, s := range stones {
		if _, ok := set[s]; ok {
			count++
		}
	}

	return count
}