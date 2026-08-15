package main

import "fmt"

func main() {
	digits := "23"
	result := letterCombinations(digits)
	fmt.Println(result)
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	alphabet := map[byte][]string{}
	alphabet['2'] = []string{"a", "b", "c"}
	alphabet['3'] = []string{"d", "e", "f"}
	alphabet['4'] = []string{"g", "h", "i"}
	alphabet['5'] = []string{"j", "k", "l"}
	alphabet['6'] = []string{"m", "n", "o"}
	alphabet['7'] = []string{"p", "q", "r", "s"}
	alphabet['8'] = []string{"t", "u", "v"}
	alphabet['9'] = []string{"w", "x", "y", "z"}
	result := []string{}
	letterCombinationsAdd(digits, "", 0, alphabet, &result)
	return result
}

func letterCombinationsAdd(digits, str string, ind int, alphabet map[byte][]string, res *[]string) {
	if ind == len(digits) {
		*res = append(*res, str)
		return
	}

	for _, l := range alphabet[digits[ind]] {
		str += l
		letterCombinationsAdd(digits, str, ind+1, alphabet, res)
		str = str[:len(str)-1]
	}
}