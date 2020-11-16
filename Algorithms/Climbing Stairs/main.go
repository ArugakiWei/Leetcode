package main

import "fmt"

func main() {
	fmt.Println(climbStairs(7))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	cur, pre1, pre2 := 0, 1, 2

	for i := 3; i <= n; i++ {
		cur = pre1 + pre2
		pre1 = pre2
		pre2 = cur
	}

	return cur
}
