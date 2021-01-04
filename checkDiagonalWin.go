package main

func checkDiagonalWin(board [3][3]string) bool {
	isWin := (board[0][0] == board[1][1] && board[1][1] == board[2][2] ||
		board[2][0] == board[1][1] && board[1][1] == board[0][2])

	return isWin
}
