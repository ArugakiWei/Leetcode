package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 2, 1, 3, 2, 2}))
	fmt.Println(numberOfPairs([]int{1, 1}))
}

func numberOfPairs(nums []int) []int {
	ans1, ans2, m := 0, 0, make(map[int]int)
	for _, n := range nums {
		m[n]++
		if m[n] == 2 {
			ans1++
			m[n] = 0
		}
	}
	for _, c := range m {
		if c != 0 {
			ans2++
		}
	}
	return []int{ans1, ans2}
}
