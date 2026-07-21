package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	result := evalRPN(tokens)
	fmt.Println(result)
}

func evalRPN(tokens []string) int {
    nums := Stack{}
	
	for _, token := range tokens {
		switch {
		case token == "+":
			val2 := nums.Pop()
			val1 := nums.Pop()
			nums.Push(val1 + val2)
		case token == "-":
			val2 := nums.Pop()
			val1 := nums.Pop()
			nums.Push(val1 - val2)
		case token == "*":
			val2 := nums.Pop()
			val1 := nums.Pop()
			nums.Push(val1 * val2)
		case token == "/":
			val2 := nums.Pop()
			val1 := nums.Pop()
			nums.Push(val1 / val2)
		default:
			val, _ := strconv.Atoi(token)
			nums.Push(val)
		}
	}

	return nums.Pop()
}

type Stack struct {
	slice []int
}

func (s *Stack) Push(val int) {
	s.slice = append(s.slice, val)
}

func (s *Stack) Peek() int {
	return s.slice[len(s.slice) - 1]
}

func (s *Stack) Pop() int {
	val := s.slice[len(s.slice) - 1]
	s.slice = s.slice[:len(s.slice) - 1]
	return val
}

func (s *Stack) Empty() bool {
	return len(s.slice) == 0
}
