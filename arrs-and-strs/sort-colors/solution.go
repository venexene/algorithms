package main

import "fmt"

func main() {
	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	fmt.Println(colors)
}

func sortColors(nums []int)  {
    low := 0
	high := len(nums)-1
	cur := 0
	for cur <= high {
		switch nums[cur] {
		case 0:
			nums[low], nums[cur] = nums[cur], nums[low]
			low++
			cur++
		case 1:
			cur++
		case 2:
			nums[high], nums[cur] = nums[cur], nums[high]
			high--
		}
	}
}