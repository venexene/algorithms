package main

import "fmt"

func main() {
	nums := []int{3,3}
	target := 6
	result := twoSumHash(nums, target)
	fmt.Println(result)
}

func twoSumBrutforce(nums []int, target int) []int {
    for i, num1 := range nums {
		for j, num2 := range nums {
			if i != j && num1 + num2 == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumTwoPointers(nums []int, target int) []int {
    i := 0
	j := len(nums) - 1
	vals := sortWithIndexes(nums)
	for i < j {
		if vals[i][0] + vals[j][0] == target {
			return []int{vals[i][1], vals[j][1]}
		} else if vals[i][0] + vals[j][0]  < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

func twoSumHash(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		if j, ok := seen[target - num]; ok {
			return []int{i, j}
		}
		seen[num] = i
	}

	return nil
}

func sortWithIndexes(nums []int) [][]int {
	vals := [][]int{}
	for i, num := range nums {
		vals = append(vals, []int{num, i})
	}

	quickSort(vals, 0, len(vals)-1)
	
	return vals
}

func quickSort(vals [][]int, low int, high int) {
	if low < high {
		pi := partition(vals, low, high)

		quickSort(vals, low, pi - 1)
		quickSort(vals, pi + 1, high)
	}
}

func partition(vals [][]int, low int, high int) int {
	pivot := vals[high][0]
	
	i := low - 1
	for j := low; j <= high; j++ {
		if vals[j][0] < pivot {
			i++
			vals[i], vals[j] = vals[j], vals[i]
		}
	}


	vals[i + 1], vals[high] = vals[high], vals[i + 1]
	return i + 1
}