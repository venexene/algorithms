package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	res := subarraySum(nums, 1)
	fmt.Println(res)
}

func subarraySum(nums []int, k int) int {
	mp := map[int]int{0: 1}
	res := 0
	cur := 0
	for _, num := range nums {
		cur += num
		res += mp[cur - k]
		mp[cur]++
	}

	return res
}