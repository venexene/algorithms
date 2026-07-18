package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequentBucket(nums, k)
	fmt.Println(result)
}

func topKFrequentHeap(nums []int, k int) []int {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}

	hp := NewTopKHeap(mp, k)
	for num := range mp {
		hp.Add(num)
	}

	return hp.Get()
}

type FreqHeap struct {
	vals []int
	freq map[int]int
} 

func NewFreqHeap(freqMap map[int]int) *FreqHeap {
	return &FreqHeap{
		vals : []int{},
		freq: freqMap,
	}
}

func (h FreqHeap) Len() int { return len(h.vals) }

func (h FreqHeap) Less(i, j int) bool { return h.freq[h.vals[i]] < h.freq[h.vals[j]] }

func (h FreqHeap) Swap(i, j int) { h.vals[i], h.vals[j] = h.vals[j], h.vals[i] }

func (h *FreqHeap) Push(x any) {
	h.vals = append(h.vals, x.(int))
}

func (h *FreqHeap) Pop() any {
	old := h.vals
	n := len(old)
	x := old[n-1]
	h.vals = old[0 : n-1]
	return x
}

type TopKHeap struct {
	heap *FreqHeap
	k int
}

func NewTopKHeap(freqMap map[int]int, k int) *TopKHeap {
	h := NewFreqHeap(freqMap)
	heap.Init(h)
	return &TopKHeap{
		heap: h,
		k:k,
	}
}

func (h *TopKHeap) Add(x int) {
	heap.Push(h.heap, x)
	if h.heap.Len() > h.k {
		heap.Pop(h.heap)
	}
}

func (h * TopKHeap) Get() []int {
	return h.heap.vals
}

func topKFrequentBucket(nums []int, k int) []int {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}
	
	buckets := make([][]int, len(nums) + 1)
	for num, val := range mp {
		buckets[val] = append(buckets[val], num)
	}

	res := []int{}
	for i := len(buckets) - 1; len(res) < k; i-- {
		res = append(res, buckets[i]...)
	}

	return res
}