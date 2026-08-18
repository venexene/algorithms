package main

import "fmt"

func main() {
	node1 := &ListNode{}
	fmt.Println(node1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	set := map[*ListNode]struct{}{}

	for head != nil {
		if _, ok := set[head]; ok {
			return true
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return false
}

func hasCycleFloyd(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil {
		slow = slow.Next

		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next

		if fast == slow {
			return true
		}
	}
	return false
}
