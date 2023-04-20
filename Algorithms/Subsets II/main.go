package main

import "sort"

func main() {

}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)

	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)
	backtrack(nums, &track, &res, 0)

	return res
}

func backtrack(nums []int, track *[]int, res *[][]int, start int) {
	// 前序位置，每个节点的值都是一个子集
	tmp := make([]int, len(*track))
	copy(tmp, *track)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		// 剪枝逻辑，值相同的相邻树枝，只遍历第一条
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*track = append(*track, nums[i])
		backtrack(nums, track, res, i+1)
		*track = (*track)[:len(*track)-1]
	}
}
