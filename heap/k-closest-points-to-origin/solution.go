package main

import (
	"fmt"
	"container/heap"
)


func main() {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	result := kClosest(points, k)
	fmt.Println(result)
}

type PointHeap struct {
	points [][]int
	k int
}

func sqrDist(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func (h PointHeap) Len() int { return len(h.points) }

func (h PointHeap) Less(i, j int) bool {
	 return sqrDist(h.points[i]) > sqrDist(h.points[j]) 
}

func (h PointHeap) Swap(i, j int) { 
	h.points[i], h.points[j] = h.points[j], h.points[i]
}

func (h *PointHeap) Push(x any) {
	h.points = append(h.points, x.([]int))
}

func (h *PointHeap) Pop() any {
	old := h.points
	n := len(old)
	x := old[n-1]
	h.points = old[0:n-1]
	return x
}

func (h *PointHeap) Add(x []int) {
	heap.Push(h, x)
	if len(h.points) > h.k {
		heap.Pop(h)
	}
}


func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{
		points: [][]int{},
		k: k,
	}

    for _, p := range points {
		h.Add(p)
	}

	return h.points
}