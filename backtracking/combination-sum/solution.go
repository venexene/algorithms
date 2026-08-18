package main

import (
	"fmt"
	"slices"
)

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	result := combinationSum(candidates, target)
	fmt.Println(result)
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	slices.Sort(candidates)
	combinationSumAdd(candidates, target, 0, []int{}, &result)
	return result
}

func combinationSumAdd(nums []int, target, start int, set []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, set...))
		return
	}
	if target < 0 {
		return
	}

	for i := start; i < len(nums); i++ {
		if nums[i] > target {
			break
		}
		set = append(set, nums[i])
		combinationSumAdd(nums, target-nums[i], i, set, res)
		set = set[:len(set)-1]
	}
}
