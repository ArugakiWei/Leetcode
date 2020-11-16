package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res, cur, pre := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(nums[i], pre+nums[i])
		pre = cur
		res = max(res, cur)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
