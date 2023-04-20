package main

import "sort"

func main() {

}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))
	var backtrack func(nums []int)

	backtrack = func(nums []int) {
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			backtrack(nums)
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	// 首先排序这个数组
	sort.Ints(nums)
	backtrack(nums)
	return res
}
