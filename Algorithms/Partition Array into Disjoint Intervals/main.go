package main

import "fmt"

func main() {
	fmt.Println(partitionDisjoint([]int{5, 0, 3, 8, 6}))
	fmt.Println(partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
	fmt.Println(partitionDisjoint([]int{1, 1}))
	fmt.Println(partitionDisjoint([]int{90, 47, 69, 10, 43, 92, 31, 73, 61, 97}))
	fmt.Println(partitionDisjoint([]int{6, 0, 8, 30, 37, 6, 75, 98, 39, 90, 63, 74, 52, 92, 64}))
	fmt.Println(partitionDisjoint([]int{29, 33, 6, 4, 42, 0, 10, 22, 62, 16, 46, 75, 100, 67, 70, 74, 87, 69, 73, 88}))
}

func partitionDisjoint(nums []int) int {
	min, minIndex := nums[0], 0
	for i, v := range nums {
		if v < min {
			minIndex, min = i, v
		}
	}
	maxLeft := nums[0]
	for i := 0; i < minIndex; i++ {
		if maxLeft < nums[i] {
			maxLeft = nums[i]
		}
	}
	lastMax, lastMaxIndex := maxLeft, 0
	for i := minIndex; i < len(nums); i++ {
		if nums[i] > lastMax {
			lastMax, lastMaxIndex = nums[i], i
		}
		if minIndex > lastMaxIndex {
			maxLeft = lastMax
		}
		if nums[i] < maxLeft {
			minIndex = i
		}
	}
	return minIndex + 1
}
