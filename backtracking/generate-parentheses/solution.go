package main

import "fmt"

func main() {
	n := 3
	result := generateParenthesis(n)
	fmt.Println(result)
}

func generateParenthesis(n int) []string {
    result := []string{}
	generateParenthesisAdd(n, 0, 0, "", &result)
	return result
}

func generateParenthesisAdd(n, open, close int, str string, res *[]string) {
	if close > open || open > n {
		return
	}

	if open + close == n * 2 {
		*res = append(*res, str)
		return
	}

	str += "("
	generateParenthesisAdd(n, open+1, close, str, res)
	str = str[:len(str)-1]
	str += ")"
	generateParenthesisAdd(n, open, close+1, str, res)
}