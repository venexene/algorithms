package main

import (
	"container/heap"
)

func main() {

}

type IntHeap struct {
	nums []int
	k    int
}

func (h IntHeap) Len() int           { return len(h.nums) }
func (h IntHeap) Less(i, j int) bool { return h.nums[i] < h.nums[j] }
func (h IntHeap) Swap(i, j int)      { h.nums[i], h.nums[j] = h.nums[j], h.nums[i] }

func (h *IntHeap) Push(x any) {
	h.nums = append(h.nums, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.nums
	n := len(old)
	x := old[n-1]
	h.nums = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() any {
	if len(h.nums) == 0 {
		return nil
	}
	return h.nums[0]
}

func (h *IntHeap) Add(x int) {
	heap.Push(h, x)
	if len(h.nums) > h.k {
		heap.Pop(h)
	}
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}

	h := IntHeap{
		nums: []int{},
		k:    k,
	}

	for _, n := range nums {
		h.Add(n)
	}

	return h.Peek().(int)
}
