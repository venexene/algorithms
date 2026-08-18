package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println(result)
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	subsetsStartIndex(nums, 0, []int{}, &result)
	return result
}

func subsetsIncludeExclude(nums []int, i int, set []int, res *[][]int) {
	if i == len(nums) {
		*res = append(*res, append([]int{}, set...))
		return
	}
	set = append(set, nums[i])
	subsetsIncludeExclude(nums, i+1, set, res)
	set = set[:len(set)-1]
	subsetsIncludeExclude(nums, i+1, set, res)
}

func subsetsStartIndex(nums []int, start int, set []int, res *[][]int) {
	*res = append(*res, append([]int{}, set...))
	for i := start; i < len(nums); i++ {
		set = append(set, nums[i])
		subsetsStartIndex(nums, i+1, set, res)
		set = set[:len(set)-1]
	}
}
