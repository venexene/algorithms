package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	result := summaryRanges(nums)
	fmt.Println(result)
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	res := []string{}
	for i := 0; i < len(nums); {
		start := nums[i]
		j := i
		for j+1 < len(nums) && nums[j+1] == nums[j]+1 {
			j++
		}
		if start == nums[j] {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[j]))
		}
		i = j + 1
	}

	return res
}