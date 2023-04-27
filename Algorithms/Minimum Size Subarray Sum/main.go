package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

func minSubArrayLen(target int, nums []int) int {
	sum, ans := 0, math.MaxInt
	left, right := 0, 0
	for right < len(nums) {
		sum += nums[right]
		right++
		for sum >= target {
			ans = min(ans, right-left)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
