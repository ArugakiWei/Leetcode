package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minElements([]int{1, -1, 1}, 3, -4))
	fmt.Println(minElements([]int{1, -10, 9, 1}, 100, 0))
}

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := goal - sum
	return ceil(abs(diff), abs(limit))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ceil(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}
