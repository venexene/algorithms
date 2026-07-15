package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicateHash(nums)
	fmt.Println(result)	
}

func containsDuplicateHash(nums []int) bool {
    hashMap := map[int]struct{}{}
	
	for _, num := range nums {
		if _, ok := hashMap[num]; ok {
			return true
		}
		hashMap[num] = struct{}{}
	}

	return false
}

func containsDuplicateSort(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i - 1] == nums[i] {
			return true
		}
	}
	return false
}