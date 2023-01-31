package main

import "fmt"

func main() {
	fmt.Println(checkXMatrix([][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}))
}

func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || j == n-1-i {
				if grid[i][j] == 0 {
					return false
				}
				continue
			}
			if grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
