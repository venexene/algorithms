package main

import "fmt"

func main() {
	s := "["
	result := isValid(s)
	fmt.Println(result)
}

func isValid(s string) bool {
	brackets := map[rune]rune{'(':')', '{':'}', '[':']'}

	stack := make([]rune, 0, len(s))
	for _, r := range s {
		if v, ok := brackets[r]; ok {
			stack = append(stack, v)
		} else if len(stack) > 0 && r == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}