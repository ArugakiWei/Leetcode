package main

import (
	"fmt"
)

func main() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7}))
	fmt.Println(prevPermOpt1([]int{3, 1, 1, 3}))
}

func prevPermOpt1(arr []int) []int {
	var i int
	for i = len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			break
		}
	}
	if i == -1 {
		return arr
	}
	var max, maxIndex int
	for j := i + 1; j < len(arr); j++ {
		if arr[j] > max && arr[j] < arr[i] {
			max, maxIndex = arr[j], j
		}
	}
	arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
	return arr
}
