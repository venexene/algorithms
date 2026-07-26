package main

func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseListIter(head *ListNode) *ListNode {
	var prev *ListNode = nil
	node := head
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

func reverseListRecur(head *ListNode) *ListNode {
	var prev *ListNode = nil
	node := head

	newHead := swapNodes(node, prev)

	return newHead
}

func swapNodes(node, prev *ListNode) *ListNode {
	if node == nil {
		return prev
	}
	next := node.Next
	node.Next = prev
	prev = node
	node = next
	return swapNodes(node, prev)
}