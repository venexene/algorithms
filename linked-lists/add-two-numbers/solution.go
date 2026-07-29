package main

import "fmt"

func main() {
	fmt.Println((2 + 5) % 10)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbersFirst(l1 *ListNode, l2 *ListNode) *ListNode {
    pred := 0
	sum := 0
	num1 := 0
	num2 := 0
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil || l2 != nil || pred != 0 {
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		} else {
			num1 = 0
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		} else {
			num2 = 0
		}

		sum = num1 + num2  + pred
		pred = sum / 10

		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}

	return tail
}