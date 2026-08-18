package main

import (
	"container/heap"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	h := NodeHeap{}
	for _, n := range lists {
		if n != nil {
			h = append(h, n)
		}
	}
	heap.Init(&h)
	for h.Len() > 0 {
		minNode := heap.Pop(&h).(*ListNode)
		node.Next = minNode
		node = node.Next
		minNode = minNode.Next
		if minNode != nil {
			heap.Push(&h, minNode)
		}
	}
	return dummy.Next
}

func mergeKListsBrutforce(lists []*ListNode) *ListNode {
	c := 0
	dummy := &ListNode{}
	node := dummy
	for {
		c = 0
		minNodeInd := 0
		for i, n := range lists {
			if n == nil {
				c++
				continue
			}
			if lists[minNodeInd] == nil || n.Val < lists[minNodeInd].Val {
				minNodeInd = i
			}
		}
		if c == len(lists) {
			break
		}
		node.Next = lists[minNodeInd]
		node = node.Next
		lists[minNodeInd] = lists[minNodeInd].Next
	}
	return dummy.Next
}
