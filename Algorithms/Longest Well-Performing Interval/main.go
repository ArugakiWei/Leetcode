package main

import "fmt"

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
	fmt.Println(longestWPI([]int{6, 6, 6}))
}

func longestWPI(hours []int) int {
	m := make(map[int]int)
	for _, h := range hours {
		if h > 8 {
			m[1]++
		} else {
			m[0]++
		}
	}

	ans := 0
	for i := 0; i < len(hours); i++ {
		lt8, gt8 := m[0], m[1]
		for j := len(hours) - 1; j >= i; j-- {
			if lt8 < gt8 {
				ans = max(ans, j-i+1)
				break
			}
			if hours[j] > 8 {
				gt8--
			} else {
				lt8--
			}
		}
		if hours[i] > 8 {
			m[1]--
		} else {
			m[0]--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
