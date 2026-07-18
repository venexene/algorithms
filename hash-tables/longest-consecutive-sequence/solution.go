package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(nums)
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 1
    set := map[int]struct{}{}

	for _, num := range nums {
		set[num] = struct{}{}
	}

	for num := range set {
		if _, ok := set[num-1]; ok {
			continue
		}

		count := 1
		cur := num
		for  {
			if _, ok := set[cur + 1]; ok {
				count++
				cur++
			} else {
				break
			}
		}
		res = max(res, count)
	}

	return res
}