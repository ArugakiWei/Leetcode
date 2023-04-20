package main

import "fmt"

func main() {
	fmt.Println(getMaximumGold([][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	}))
	fmt.Println(getMaximumGold([][]int{
		{1, 6, 1},
		{5, 8, 7},
		{1, 9, 1},
	}))
	fmt.Println(getMaximumGold([][]int{
		{0, 0, 34, 0, 5, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 21, 0, 0, 0, 0, 0},
		{0, 18, 0, 0, 8, 0, 0, 0, 4, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{15, 0, 0, 0, 0, 22, 0, 0, 0, 21},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 0, 0, 0, 0, 38, 0},
	}))
}

func getMaximumGold(grid [][]int) int {
	var ans, curMax int
	used := make([][]bool, len(grid))
	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, len(grid[i]))
	}
	var path []int
	var backtrack func(row, col int)
	move := func(row, col int) {
		if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row]) && !used[row][col] && grid[row][col] != 0 {
			curMax += grid[row][col]
			used[row][col] = true
			path = append(path, grid[row][col])
			backtrack(row, col)
			curMax -= grid[row][col]
			used[row][col] = false
			path = path[:len(path)-1]
		}
	}
	backtrack = func(row, col int) {
		ans = max(ans, curMax)
		move(row-1, col)
		move(row+1, col)
		move(row, col-1)
		move(row, col+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				curMax += grid[i][j]
				used[i][j] = true
				backtrack(i, j)
				curMax -= grid[i][j]
				used[i][j] = false
			}
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
