package main

import "fmt"

func main() {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}

func isValidSudoku(board [][]byte) bool {
	var hm, vm, zm map[byte]struct{}
	for i := 0; i < len(board); i++ {
		hm = make(map[byte]struct{})
		vm = make(map[byte]struct{})
		for j := 0; j < len(board[i]); j++ {
			// 横
			if board[i][j] != '.' {
				if _, ok := hm[board[i][j]]; ok {
					return false
				} else {
					hm[board[i][j]] = struct{}{}
				}
			}
			// 竖
			if board[j][i] != '.' {
				if _, ok := vm[board[j][i]]; ok {
					return false
				} else {
					vm[board[j][i]] = struct{}{}
				}
			}
			// 3*3
			if i%3 == 0 && j%3 == 0 {
				zm = make(map[byte]struct{})
				for z := 0; z < 3; z++ {
					for x := 0; x < 3; x++ {
						if board[i+z][j+x] != '.' {
							if _, ok := zm[board[i+z][j+x]]; ok {
								return false
							} else {
								zm[board[i+z][j+x]] = struct{}{}
							}
						}
					}
				}
			}
		}
	}
	return true
}
