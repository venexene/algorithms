package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersection(nums1, nums2)
	fmt.Println(result)
}

func intersection(nums1 []int, nums2 []int) []int {
    set := map[int]struct{}{}
	inter := []int{}

	for _, num := range nums1 {
		set[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			inter = append(inter, num)
			delete(set, num)
		}
	}

	return inter
}