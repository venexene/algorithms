package main

import "fmt"

func main() {
	nums := []int{2, 3, 0, 1, 4}
	result := jump(nums)
	fmt.Println(result)
}

func jump(nums []int) int {
	mx := nums[0]
	mx2 := nums[0]
	res := 1

	if len(nums) == 1 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if mx >= len(nums)-1 {
			return res
		}

		if mx < i {
			res++
			mx = mx2
		}
		mx2 = max(mx2, i+nums[i])
	}

	return res
}
