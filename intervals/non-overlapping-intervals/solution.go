package main

import (
	"cmp"
	"slices"
)

func main() {

}

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
    
	prevEnd := intervals[0][1]
	res := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prevEnd {
			res++
			prevEnd = min(prevEnd, intervals[i][1])
			continue	
		}
		prevEnd = intervals[i][1]
	}

	return res
}