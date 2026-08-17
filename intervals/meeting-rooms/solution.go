package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	intervals := []Interval{{0, 30}, {5, 10}, {15, 20}}
	result := canAttendMeetings(intervals)
	fmt.Println(result)
}

type Interval struct {
	start int
	end   int
}

func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})
	
	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < intervals[i-1].end {
			return false
		}
	}

	return true
}