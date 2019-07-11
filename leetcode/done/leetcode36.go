package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidSudoku([][]byte{}))
}
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		m := make(map[byte]int)
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
			} else {
				if _, ok := m[board[i][j]]; ok == false {
					m[board[i][j]]++
				} else {
					return false
				}
			}
		}
	}
	for j := 0; j < len(board); j++ {
		m := make(map[byte]int)
		for i := 0; i < len(board[j]); i++ {
			if board[i][j] == '.' {
			} else {
				if _, ok := m[board[i][j]]; ok == false {
					m[board[i][j]]++
				} else {
					return false
				}
			}

		}
	}
	for r := 0; r < len(board)/3; r++ {
		for k := 0; k < len(board)/3; k++ {
			m := make(map[byte]int)
			for i := r * 3; i < (r+1)*3; i++ {
				for j := k * 3; j < (k+1)*3; j++ {
					if board[i][j] == '.' {
					} else {
						if _, ok := m[board[i][j]]; ok == false {
							m[board[i][j]]++
						} else {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
