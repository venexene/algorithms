package main

import (
	"fmt"
	"strconv"
)

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	result := calPoints(ops)
	fmt.Println(result)
}

func calPoints(operations []string) int {
    records := &Stack{}

	for _, op := range operations {
		switch {
		case op == "+":
			val1 := records.Pop()
			val2 := records.Peek()
			records.Push(val1)
			records.Push(val1 + val2)
		case op == "D":
			val := records.Peek()
			records.Push(val * 2)
		case op == "C":
			records.Pop() 
		default:
			val, _ := strconv.Atoi(op)
			records.Push(val)
		}
	}
	
	sum := 0
	for !records.Empty() {
		sum += records.Pop()
	}

	return sum
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