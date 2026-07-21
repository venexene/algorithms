package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(temperatures)
	fmt.Println(result)
}

func dailyTemperatures(temperatures []int) []int {
    results := make([]int, len(temperatures))
	dayStack := Stack{}

	dayStack.Push(0)
	for i:= 1; i < len(temperatures); i++ {
		for !dayStack.Empty() && temperatures[i] > temperatures[dayStack.Peek()]  {
			p := dayStack.Peek()
			results[p] = i - p
			dayStack.Pop()
		}
		dayStack.Push(i)
	}

	return results
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