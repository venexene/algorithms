package main

func main() {
	myQueue := Constructor();
	myQueue.Push(1); // queue is: [1]
	myQueue.Push(2); // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Peek(); // return 1
	myQueue.Pop(); // return 1, queue is [2]
	myQueue.Empty(); // return false
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

type MyQueue struct {
	funcStack *Stack
	mainStack *Stack
}


func Constructor() MyQueue {
    return MyQueue{
		funcStack: &Stack{},
		mainStack: &Stack{},
	}
}


func (this *MyQueue) Push(x int)  {
	this.funcStack.Push(x)
}


func (this *MyQueue) Pop() int {
    if this.mainStack.Empty() {
		for !this.funcStack.Empty() {
			this.mainStack.Push(this.funcStack.Pop())
		}
	}
	return this.mainStack.Pop()
}


func (this *MyQueue) Peek() int {
    for !this.funcStack.Empty() {
		this.mainStack.Push(this.funcStack.Pop())
	}
	return this.mainStack.Peek()
}


func (this *MyQueue) Empty() bool {
    return this.mainStack.Empty() && this.funcStack.Empty()
}