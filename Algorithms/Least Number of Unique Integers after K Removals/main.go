package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1))
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	countArr := make([]int, 0, len(m))
	for _, c := range m {
		countArr = append(countArr, c)
	}
	sort.Ints(countArr)

	var result int
	for i := 0; i < len(countArr); i++ {
		if k >= countArr[i] {
			k = k - countArr[i]
			result++
		}
		if k < countArr[i] {
			return len(countArr) - result
		}
	}
	return len(countArr) - result
}
