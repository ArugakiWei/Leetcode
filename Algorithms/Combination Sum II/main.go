package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans, path := make([][]int, 0, 0), make([]int, 0, 0)
	trackSum := 0
	var backtrack func(nums []int, start, target int)
	backtrack = func(nums []int, start, target int) {
		if trackSum == target {
			t := make([]int, len(path))
			copy(t, path)
			ans = append(ans, t)
		}
		if trackSum > target {
			return
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			trackSum += nums[i]
			backtrack(nums, i+1, target)
			trackSum -= nums[i]
			path = path[:len(path)-1]
		}
	}
	backtrack(candidates, 0, target)
	return ans
}
