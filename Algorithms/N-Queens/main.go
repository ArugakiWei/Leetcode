package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(5))
}

var (
	result = make([][]string, 0)
)

func solveNQueens(n int) [][]string {
	result = make([][]string, 0)
	backtrack(generate(n), 0)
	return result
}

func generate(n int) [][]byte {
	m := make([][]byte, n)
	for i:=0; i<n; i++ {
		var mi []byte
		for j:=0; j<n; j++ {
			mi = append(mi, '.')
		}
		m[i] = mi
	}
	return m
}

func isValid(m [][]byte, i, j int) bool {
	// 判断第i行是否有Q, 如果有, 返回 false
	for _, b := range m[i] {
		if b == 'Q' {
			return false
		}
	}
	// 判断第j列是否有Q, 如果有, 返回 false
	for _, col := range m {
		if col[j] == 'Q' {
			return false
		}
	}

	// 判断右上是否有Q, 如果有, 返回 false
	for ii, jj := i, j; (ii < len(m) && ii > 0) || (jj < len(m) && jj > 0); {
		ii, jj = ii+1, jj+1
		if ii > len(m)-1 || ii < 0 || jj > len(m)-1 || jj < 0 {
			break
		}
		if m[ii][jj] == 'Q' {
			return false
		}
	}
	// 判断左上是否有Q, 如果有, 返回 false
	for ii, jj := i, j; (ii < len(m) && ii > 0) || (jj < len(m) && jj > 0); {
		ii, jj = ii-1, jj+1
		if ii > len(m)-1 || ii < 0 || jj > len(m)-1 || jj < 0 {
			break
		}
		if m[ii][jj] == 'Q' {
			return false
		}
	}
	// 判断右下是否有Q, 如果有, 返回 false
	for ii, jj := i, j; (ii < len(m) && ii > 0) || (jj < len(m) && jj > 0); {
		ii, jj = ii+1, jj-1
		if ii > len(m)-1 || ii < 0 || jj > len(m)-1 || jj < 0 {
			break
		}
		if m[ii][jj] == 'Q' {
			return false
		}
	}
	// 判断左下是否有Q, 如果有, 返回 false
	for ii, jj := i, j; (ii < len(m) && ii > 0) || (jj < len(m) && jj > 0); {
		ii, jj = ii-1, jj-1
		if ii > len(m)-1 || ii < 0 || jj > len(m)-1 || jj < 0 {
			break
		}
		if m[ii][jj] == 'Q' {
			return false
		}
	}
	return true
}

func backtrack(board [][]byte, row int) {
	if row == len(board) {
		tmp := make([]string, 0, row)
		for _, v := range board {
			tmp = append(tmp, string(v))
		}

		result = append(result, tmp)
		return
	}

	for col:=0; col<len(board[row]); col++ {
		if !isValid(board, row, col) {
			continue
		}

		board[row][col] = 'Q'
		backtrack(board, row + 1)
		board[row][col] = '.'
	}
}