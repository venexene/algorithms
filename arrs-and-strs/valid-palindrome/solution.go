package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := " "
	result := isPalindromeTwoPointers(s)
	fmt.Println(result)
}

func isPalindromeTwoPointers(s string) bool {
	i := 0
	j := len(s) - 1
	runes := []rune(s)
	for i < j {
		for i < j && !isAlphanumeric(runes[i]) {
			i++
		}

		for i < j && !isAlphanumeric(runes[j]) {
			j--
		}

		if i >= j {
			break
		}

		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isPalindromeFilter(s string) bool {
	runes := []rune(s)

	ln := len(runes)

	var fwd strings.Builder
	var rev strings.Builder

	fwd.Grow(ln)
	rev.Grow(ln)

	for i := 0; i < ln; i++ {
		if isAlphanumeric(runes[i]) {
			fwd.WriteRune(unicode.ToLower(runes[i]))
		}
		if isAlphanumeric(runes[ln - 1 - i]) {
			rev.WriteRune(unicode.ToLower(runes[ln - 1 - i]))
		}
	}

	return fwd.String() == rev.String()
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}