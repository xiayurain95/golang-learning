package main

func main() {

}
func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	boardTemp := make([][]byte, len(board))
	for i := 0; i < len(board); i++ {
		boardTemp[i] = make([]byte, len(board[0]))
	}
	copy(boardTemp, board)
	ch := make(chan point, 1000)
	filterBoarder(board, ch)

	for len(ch) != 0 {

	}
}
func filterBoarder(board [][]byte, ch chan point) {
	for i, j, k := 0, len(board)-1, 0; k < len(board[0]); k++ {
		if board[i][k] == 'O' {
			ch <- point{i, k}
		}
		if board[j][k] == 'O' {
			ch <- point{j, k}
		}
	}
	for i, j, k := 0, len(board[0]), 1; k < len(board)-1; k++ {
		if board[k][i] == 'O' {
			ch <- point{k, i}
		}
		if board[k][j] == 'O' {
			ch <- point{k, j}
		}
	}
}

type point struct {
	i, j int
}
