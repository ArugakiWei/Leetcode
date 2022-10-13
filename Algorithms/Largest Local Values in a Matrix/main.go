package main

import "fmt"

func main() {
	fmt.Println(largestLocal([][]int{
		{9, 9, 8, 1},
		{5, 6, 2, 6},
		{8, 2, 6, 4},
		{6, 2, 2, 2},
	}))
	fmt.Println(largestLocal([][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}))
}

func largestLocal(grid [][]int) [][]int {
	result := make([][]int, len(grid)-2)
	for i := 0; i < len(grid)-2; i++ {
		result[i] = make([]int, len(grid)-2)
		for j := 0; j < len(grid[i])-2; j++ {
			max := grid[i][j]
			for x := 0; x < 3; x++ {
				for z := 0; z < 3; z++ {
					if grid[i+x][j+z] > max {
						max = grid[i+x][j+z]
					}
				}
			}
			result[i][j] = max
		}
	}
	return result
}
