package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
	// fmt.Println(maximalSquare1([][]byte{
	// 	{'0', '0', '0', '1'},
	// 	{'1', '1', '0', '1'},
	// 	{'1', '1', '1', '1'},
	// 	{'0', '1', '1', '1'},
	// 	{'0', '1', '1', '1'},
	// }))
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m := make([][]int, len(matrix))
	for i, v := range matrix {
		tm := make([]int, len(v))
		if i == 0 {
			for j, vv := range v {
				if vv == '1' {
					tm[j] = 1
				}
			}
		} else {
			if v[0] == '1' {
				tm[0] = 1
			}
		}
		m[i] = tm
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			l := m[i-1][j-1]
			if matrix[i][j] == '1' {
				m[i][j] = 1
				for ll := l; ll > 0; ll-- {
					if isSquare(matrix, i, j, ll) {
						m[i][j] = ll + 1
						break
					}
				}
			}
		}
	}

	var max int
	for _, v := range m {
		for _, vv := range v {
			if vv > max {
				max = vv
			}
		}
	}
	return max * max
}

func isSquare(matrix [][]byte, i, j, l int) bool {
	if i-l < 0 || j-l < 0 {
		return false
	}
	for ; l > 0; l-- {
		if matrix[i-l][j] != '1' || matrix[i][j-l] != '1' {
			return false
		}
	}
	return true
}
