package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}

func searchRange(nums []int, target int) []int {
	notFound := []int{-1, -1}
	result := []int{0, 0}

	if len(nums) == 0 {
		return notFound
	}

	// 左边界
	left, mid, right := 0, 0, len(nums) - 1
	for left <= right {
		mid = left + (right - left) / 2
		if nums[mid] == target {
			right = mid - 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return notFound
	}
	result[0] = left

	// 右边界
	left, mid, right = 0, 0, len(nums) - 1
	for left <= right {
		mid = left + (right - left) / 2
		if nums[mid] == target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return notFound
	}
	result[1] = right

	return result
}