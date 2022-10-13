package main

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	board = [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}
	solveSudoku(board)
}

var (
	result [][]byte
)

func solveSudoku(board [][]byte) {
	result = make([][]byte, len(board))
	fill(board, 0, 0)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			board[i][j] = result[i][j]
		}
	}
}

func fill(board [][]byte, row, col int) bool {
	if row == len(board)-1 && col == len(board[row]) {
		for i := 0; i < len(board); i++ {
			result[i] = make([]byte, len(board[i]))
			for j := 0; j < len(board[i]); j++ {
				result[i][j] = board[i][j]
			}
		}
		return true
	}
	if col == len(board[row]) {
		return fill(board, row+1, 0)
	}
	if board[row][col] != '.' {
		return fill(board, row, col+1)
	}

	var i byte
	for i = '1'; i <= '9'; i++ {
		if !isValidSudoku(board, row, col, i) {
			continue
		}
		board[row][col] = i
		if fill(board, row, col+1) {
			break
		}
		board[row][col] = '.'
	}
	return false
}

func isValidSudoku(board [][]byte, row, col int, x byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == x {
			return false
		}
		if board[i][col] == x {
			return false
		}
		// 找左上角第一点
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == x {
			return false
		}
	}
	return true
}
