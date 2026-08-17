package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	intervals := [][]int{{1, 3}}
	result := merge(intervals)
	fmt.Println(result)
}

func merge(intervals [][]int) [][]int {
    slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	merged := make([][]int, 0, len(intervals))

	i := 0
	for i < len(intervals) {
		start := intervals[i][0]
		end := intervals[i][1]
		j := i
		for j+1 < len(intervals) && intervals[j+1][0] <= end {
			j++
			end = max(end, intervals[j][1])
		}
		merged = append(merged, []int{start, end})
		i = j + 1
	}

	return merged
}