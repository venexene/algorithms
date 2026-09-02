package main

import (
	"cmp"
	"slices"
)

type Job struct {
	Start int
	End int
	Profit int
}

func findLastNoIntersect(jobs []Job, n int) int {
	left := 0
	right := n - 1
	result := -1
	for left <= right {
		mid := (left + right)/2
		if jobs[mid].End <= jobs[n].Start {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1  
		}
	}

	return result
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
    jobs := []Job{}
	
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, Job{startTime[i], endTime[i], profit[i]})
	}

	slices.SortFunc(jobs, func(a, b Job) int {
		return cmp.Compare(a.End, b.End)
	})

	dp := make([]int, len(jobs))
	dp[0] = jobs[0].Profit
	
	for i := 1; i < len(jobs); i++ {
		ind := findLastNoIntersect(jobs, i)
		
		prev := 0
		if ind != -1 {
			prev = dp[ind]
		}

		dp[i] = max(prev + jobs[i].Profit, dp[i-1])
	}

	return dp[len(dp) - 1]
}