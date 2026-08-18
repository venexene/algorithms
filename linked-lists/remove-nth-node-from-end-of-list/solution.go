package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	var p2 *ListNode = nil
	var prev *ListNode = nil
	c := 0

	if p1.Next == nil {
		return nil
	}

	for p1 != nil {
		if c == n {
			p2 = head
		}
		p1 = p1.Next
		c++

		if p2 != nil {
			prev = p2
			p2 = p2.Next
		}
	}

	if p2 == nil {
		head = head.Next
	} else {
		prev.Next = p2.Next
	}

	return head
}

func removeNthFromEndDummy(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow := dummy
	fast := dummy

	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
