package main

import "fmt"

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	result := insert(intervals, newInterval)
	fmt.Println(result)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	i := 0

	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		i++
	}
	fmt.Println(i)
	res = append(res, intervals[:i]...)

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)

	res = append(res, intervals[i:]...)

	return res
}