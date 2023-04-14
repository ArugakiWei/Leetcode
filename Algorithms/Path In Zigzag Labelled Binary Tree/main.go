package main

import (
	"math"
)

func main() {
	pathInZigZagTree(14)
}

func pathInZigZagTree(label int) []int {
	var n []int
	sum := 0
	for i := 0; ; i++ {
		sum += int(math.Pow(2, float64(i)))
		n = append(n, sum)
		if sum >= label {
			break
		}
	}
	ans1 := []int{label}
	p := label
	for i := len(n) - 2; i >= 0; i-- {
		c := math.Ceil(float64(p-n[i]) / 2)
		p = n[i] + 1 - int(c)
		ans1 = append(ans1, p)
	}
	ans := make([]int, 0, len(ans1))
	for i := len(ans1) - 1; i >= 0; i-- {
		ans = append(ans, ans1[i])
	}
	return ans
}
