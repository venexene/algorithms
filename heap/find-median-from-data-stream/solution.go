package main

import (
	"container/heap"
	"fmt"
)

func main() {
	
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MaxHeap) Peek() any {
	if len(h) == 0 {
		return nil
	}
	return h[0]
}


type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MinHeap) Peek() any {
	if len(h) == 0 {
		return nil
	}
	return h[0]
}

type MedianFinder struct {
    minHeap *MinHeap
	maxHeap *MaxHeap
}


func Constructor() MedianFinder {
    return MedianFinder{
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}


func (this *MedianFinder) AddNum(num int)  {
	if this.maxHeap.Len() == 0 && this.minHeap.Len() == 0 {
		heap.Push(this.maxHeap, num)
		return
	}

	if num <= this.maxHeap.Peek().(int) {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}

	if this.minHeap.Len() - this.maxHeap.Len() > 1 {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	} else if this.maxHeap.Len() - this.minHeap.Len() > 1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == this.maxHeap.Len() {
		fmt.Println("==")
		return float64(this.minHeap.Peek().(int) + this.maxHeap.Peek().(int)) / 2
	} else if this.minHeap.Len() > this.maxHeap.Len() {
		fmt.Println("min")
		return float64(this.minHeap.Peek().(int))
	} else {
		fmt.Println("max")
		return float64(this.maxHeap.Peek().(int))
	}
}
