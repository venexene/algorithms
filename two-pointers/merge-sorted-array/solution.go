package main

import "fmt"

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	n1 := m - 1
	n2 := n - 1

	for i := len(nums1) - 1; i >= 0; i-- {

		if n2 < 0 {
			break
		}

		if n1 >= 0 && nums1[n1] > nums2[n2] {
			nums1[i] = nums1[n1]
			n1--
		} else {
			nums1[i] = nums2[n2]
			n2--
		}
	}
}