package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	result := searchOptimal(nums, target)
	fmt.Println(result)
}

func searchFirst(nums []int, target int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			k = i
			break
		}
	}

	low := 0
	high := len(nums) - 1

	fmt.Println(low, high)

	for low <= high {
		mid := (low + high) / 2
		mInd := (k + mid) % len(nums)

		if nums[mInd] == target {
			return mInd
		}

		if target > nums[mInd] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func searchOptimal(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	fmt.Println(low, high)

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
