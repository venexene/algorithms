package main

import (
	"container/heap"
	"fmt"
)

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	result := lastStoneWeight(stones)
	fmt.Println(result)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := IntHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		stone1 := heap.Pop(&h).(int)
		stone2 := heap.Pop(&h).(int)
		heap.Push(&h, stone1-stone2)
	}
	return heap.Pop(&h).(int)
}
