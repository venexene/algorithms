package main

func main() {

}

type MinStack struct {
    slice []int
	minStack Stack
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(value int)  {
	if len(this.slice) == 0 ||  value <= this.minStack.Peek() {
		this.minStack.Push(value)
	}
	this.slice = append(this.slice, value)
}


func (this *MinStack) Pop()  {
	if this.Top() == this.minStack.Peek() {
		this.minStack.Pop()
	}
	this.slice = this.slice[:len(this.slice) - 1]
}


func (this *MinStack) Top() int {
	return this.slice[len(this.slice) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minStack.Peek()
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