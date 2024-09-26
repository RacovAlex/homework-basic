package main

import "sort"

func main() {}

func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	sort.Ints(nums)

	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			high = mid - 1
		default:
			low = mid + 1
		}
	}

	return -1
}
