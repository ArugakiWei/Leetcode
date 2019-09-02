package main

import (
	"fmt"
)

func smallestRangeI(A []int, K int) int {
	if len(A) == 0 || len(A) == 1 {
		return 0
	}

	max, min := A[0], A[0]
	for _, val := range A {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	if max - min < 2 * K {
		return 0;
	} else {
		return max - min - (2 * K);
	}
}

func main() {
	a := []int{0, 10}
	k := 2
	fmt.Println(smallestRangeI(a, k))

	a = []int{1,3,6}
	k = 3
	fmt.Println(smallestRangeI(a, k))

	a = []int{7,8,8,5,2}
	k = 4
	fmt.Println(smallestRangeI(a, k))

	a = []int{9,9,2,8,7}
	k = 4
	fmt.Println(smallestRangeI(a, k))
}