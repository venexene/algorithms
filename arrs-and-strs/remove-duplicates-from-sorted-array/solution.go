package main

import "fmt"

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	result := removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	pos := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i - 1] {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}