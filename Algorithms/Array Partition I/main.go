package main

import (
	"sort"
	"fmt"
)

func main() {

	fmt.Println(arrayPairSum([]int{1,4,3,2}))
}

func arrayPairSum(nums []int) int {

	sort.Ints(nums)
	sum := 0
	for i, v := range nums {
		if (i+1)%2 != 0 {
			sum += v
		}
	}
	return sum
}