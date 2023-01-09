package main

import (
	"fmt"
)

func main() {
	for i := 2; i < 100; i += 2 {
		fmt.Println(i, reinitializePermutation(i))
	}
}

func reinitializePermutation(n int) int {
	var perm []int
	for i := 0; i < n; i++ {
		perm = append(perm, i)
	}

	arr := operate(perm)
	ans := 1
	for {
		if isEqual(perm, arr) {
			break
		}
		arr = operate(arr)
		ans++
	}
	return ans
}

func isEqual(perm, arr []int) bool {
	for i, x := range perm {
		if x != arr[i] {
			return false
		}
	}
	return true
}

func operate(src []int) []int {
	dst := make([]int, len(src))
	for i, _ := range src {
		if i%2 == 0 {
			dst[i] = src[i/2]
		} else {
			dst[i] = src[len(src)/2+(i-1)/2]
		}
	}
	return dst
}
