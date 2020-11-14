package main

import "fmt"

func main() {
	fmt.Println(permute([]int{0,1}))
}

var (
	result = make([][]int, 0)
)

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	backtrack(nums, make([]int, 0, len(nums)))
	return result
}

func backtrack(nums []int, singleResult []int) {
	if len(nums) == len(singleResult) {
		tmp := make([]int, len(singleResult))
		copy(tmp, singleResult)
		result = append(result, tmp)
		return
	}

	for _, i := range nums {
		if contains(singleResult, i) {
			continue
		}

		singleResult = append(singleResult, i)
		backtrack(nums, singleResult)
		singleResult = singleResult[:len(singleResult)-1]
	}
}

func contains(content []int, value int) bool {
	for _, i := range content {
		if i == value {
			return true
		}
	}
	return false
}