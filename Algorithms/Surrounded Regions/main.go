package main

import "fmt"

func main() {
	a := [][]byte{{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O'}, {'X', 'O', 'O', 'X', 'X', 'O', 'O', 'X', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O', 'O', 'X', 'O'}, {'O', 'O', 'O', 'X', 'X', 'X', 'X', 'O', 'X', 'O', 'X', 'X', 'O', 'O', 'O', 'O', 'X', 'O', 'X', 'O'}, {'O', 'O', 'O', 'X', 'X', 'O', 'O', 'X', 'O', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'X', 'O', 'O', 'X'}, {'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O'}, {'X', 'O', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'X'}, {'O', 'O', 'O', 'X', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O', 'X', 'O', 'X', 'O', 'X', 'O', 'X'}, {'O', 'O', 'O', 'X', 'O', 'X', 'O', 'O', 'X', 'X', 'O', 'X', 'O', 'X', 'X', 'O', 'X', 'X', 'X', 'O'}, {'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X', 'X', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'X', 'O', 'X'}, {'O', 'O', 'X', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X', 'O', 'O', 'O', 'X', 'O'}, {'X', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'X', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'X'}, {'O', 'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'O', 'O', 'O'}, {'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X', 'X', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O'}, {'X', 'O', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'X'}, {'X', 'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O'}, {'O', 'X', 'X', 'O', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O'}, {'O', 'X', 'O', 'O', 'O', 'O', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X'}, {'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X', 'X', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'X', 'O', 'O', 'O'}}
	output(a)

	solve(a)
	output(a)
}

func output(board [][]byte) {
	for _, v := range board {
		fmt.Println(string(v))
	}
	fmt.Println()
}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	replace(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// a 换成 b
func replace(board [][]byte) {
	for i, _ := range board[0] {
		dfs(board, 0, i)            // 第一行
		dfs(board, len(board)-1, i) // 最后一行
	}
	for i, _ := range board {
		dfs(board, i, 0)               // 第一列
		dfs(board, i, len(board[0])-1) // 最后一列
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return
	}
	if board[i][j] != 'O' {
		return
	}

	board[i][j] = '#'

	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
}
