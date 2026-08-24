package main

import "fmt"

func main() {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	result := wordBreak(s, wordDict)
	fmt.Println(result)
}

func wordBreak(s string, wordDict []string) bool {
	dict := map[string]bool{}
	for _, w := range wordDict {
		dict[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakBacktracking(s string, wordDict []string) bool {
	dict := map[string]struct{}{}
	for _, s := range wordDict {
		dict[s] = struct{}{}
	}

	return wordBreakAdd(s, dict, 0)
}

func wordBreakAdd(s string, dict map[string]struct{}, i int) bool {
	if len(s) <= 0 {
		return true
	}

	if i > len(s) {
		return false
	}

	res := false
	res = res || wordBreakAdd(s, dict, i+1)
	if _, ok := dict[s[:i]]; ok {
		res = res || wordBreakAdd(s[i:], dict, 0)
	}
	return res
}
