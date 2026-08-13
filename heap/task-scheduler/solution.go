package main

import (
	"container/heap"
)

func main() {

}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

type Waiter struct {
	c int
	w int
}

func leastInterval(tasks []byte, n int) int {
	counter := map[byte]int{}
	for _, task := range tasks {
		counter[task]++
	}

	count := []int{}
	for _, v := range counter {
		count = append(count, v)
	} 

    h := IntHeap(count)
	heap.Init(&h)

	waiters := []Waiter{}
	res := 0
	for h.Len() > 0 || len(waiters) > 0	{
		newWaiters := waiters[:0]
		for i := range waiters {
			waiters[i].w--
			if waiters[i].w == 0 {
				heap.Push(&h, waiters[i].c)
			} else {
				newWaiters = append(newWaiters, waiters[i])
			}
		}
		waiters = newWaiters

		res++
		if h.Len() == 0 {
			continue
		}

		c := heap.Pop(&h).(int)
		c--
		if c != 0 {
			waiters = append(waiters, Waiter{c: c, w: n+1})
		}
	}

	return res
}