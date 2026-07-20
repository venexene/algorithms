package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	result := threeSumMap(nums)
	fmt.Println(result)
}

func threeSumSort(nums []int) [][]int {
    set := map[[3]int]struct{}{}
	
	sort.Ints(nums)
	for i := 0; i < len(nums) - 2; i++ {
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left] + nums[right] > target {
				right--
			} else if nums[left] + nums[right] < target {
				left++
			} else {
				set[[3]int{nums[i], nums[right], nums[left]}] = struct{}{}
				left++
				right--
			}
		}
	}


	keys := make([][]int, 0, len(set))
	for k := range set {
		vals := make([]int, 0, 3)
		for _, v := range k {
			vals = append(vals, v)
		}
		keys = append(keys, vals)
	}

	return keys 
}

func threeSumMap(nums []int) [][]int {
    set := map[[3]int]struct{}{}
	
	for i := 0; i < len(nums) - 2; i++ {
		target := -nums[i]
		
		seen := map[int]int{}
		for j := i + 1; j < len(nums); j++ {
			if k, ok := seen[target - nums[j]]; ok {
				key := [3]int{nums[i], nums[j], nums[k]}
				slices.Sort(key[:])
				set[key] = struct{}{}
			}
			seen[nums[j]] = j
		}
	}

	keys := make([][]int, 0, len(set))
	for k := range set {
		vals := make([]int, 0, 3)
		for _, v := range k {
			vals = append(vals, v)
		}
		keys = append(keys, vals)
	}

	return keys 
}

func threeSumSortNoSet(nums []int) [][]int {
	res := make([][]int, 0, len(nums))

	sort.Ints(nums)
	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left] + nums[right] > target {
				right--
			} else if nums[left] + nums[right] < target {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}


	return res
}