package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	fmt.Println(longestOnes([]int{0, 0, 0, 1}, 4))
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 0))
}

func longestOnes(nums []int, k int) int {
	ans := 0
	left, right, zeroCount := 0, 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}
		right++

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
