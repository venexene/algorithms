package main

import (
	"slices"
	"fmt"
)

func main() {
	intervals := []Interval{{0, 40}, {5, 10}, {15, 20}}
	result := minMeetingRooms(intervals)
	fmt.Println(result)
}

type Interval struct {
	start int
	end   int
}

func minMeetingRooms(intervals []Interval) int {
	starts := make([]int, 0, len(intervals))
	ends := make([]int, 0, len(intervals))
	for _, i := range intervals {
		starts = append(starts, i.start)
		ends = append(ends, i.end)

	}
	slices.Sort(starts)
	slices.Sort(ends)

	maxRooms := 0
	actRooms := 0
	j := 0
	for i := 0; i < len(starts); i++ {
		if ends[j] <= starts[i] {
			actRooms--
			j++
		}
		actRooms++
		maxRooms = max(maxRooms, actRooms)
	}

	return maxRooms
}