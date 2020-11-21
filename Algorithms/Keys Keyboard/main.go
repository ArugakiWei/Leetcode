package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSteps(180))
}

func minSteps(n int) int {
	ans, d := 0, 2
	for n > 1 {
		for n%d == 0 {
			ans += d
			n /= d
		}
		d++
	}
	return ans
}
