package main

import "fmt"

func main() {
	nums := []int{2}
	val := 3
	result := removeElement(nums, val)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	pos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}

	return pos 
}