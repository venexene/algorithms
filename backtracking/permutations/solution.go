package main

import (
	"fmt"
	"maps"
)

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println(result)
}

func permute(nums []int) [][]int {
    result := [][]int{}
	permuteSwap(nums, 0, &result)
	return result
}

func permuteSwap(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		*res = append(*res, append([]int{}, nums...))
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permuteSwap(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permuteAddSlice(nums []int, set []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, append([]int{}, set...))
		return
	}

	for i := 0; i < len(nums); i++ {
		set = append(set, nums[i])
		newNums := make([]int, len(nums)-1)
		copy(newNums[:i], nums[:i])
		copy(newNums[i:], nums[i+1:])
		permuteAddSlice(newNums, set, res)
		set = set[:len(set)-1]
	}
}

func permuteAddMap(nums []int, visited map[int]struct{}, set []int, res *[][]int) {
	if len(set) == len(nums) {
		*res = append(*res, append([]int{}, set...))
		return
	}

	if len(set) > 0 {
		visited[set[len(set)-1]] = struct{}{}
	}

	for _, n := range nums {
		if _, ok := visited[n]; !ok {
			set = append(set, n)
			permuteAddMap(nums, maps.Clone(visited) , set, res)
			set = set[:len(set)-1]
		}
	}
}