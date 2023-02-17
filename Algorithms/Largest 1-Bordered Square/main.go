package main

import "fmt"

func main() {
	fmt.Println(largest1BorderedSquare([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}))
	fmt.Println(largest1BorderedSquare([][]int{
		{0, 0, 1, 0},
	}))
	fmt.Println(largest1BorderedSquare([][]int{
		{0, 1},
		{1, 1},
	}))
}

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rs := make([][]int, m)
	for i := range rs {
		rs[i] = make([]int, n+1)
	}
	cs := make([][]int, n)
	for i := range cs {
		cs[i] = make([]int, m+1)
	}
	for i, row := range grid {
		for j, x := range row {
			rs[i][j+1] = rs[i][j] + x // 每行的前缀和
			cs[j][i+1] = cs[j][i] + x // 每列的前缀和
		}
	}
	for d := min(m, n); d > 0; d-- {
		for i := 0; i <= m-d; i++ {
			for j := 0; j <= n-d; j++ {
				if rs[i][j+d]-rs[i][j] == d && // 上边
					cs[j][i+d]-cs[j][i] == d && // 左边
					rs[i+d-1][j+d]-rs[i+d-1][j] == d && // 下边
					cs[j+d-1][i+d]-cs[j+d-1][i] == d { // 右边
					return d * d
				}
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
