package main

// Checks if on of the players won.
// Returns false if no winner true otherwise.
func checkWin(board [3][3]string) bool {

	horiz := (board[0][0] == board[0][1] && board[0][1] == board[0][2] ||
		board[1][0] == board[1][1] && board[1][1] == board[1][2] ||
		board[2][0] == board[2][1] && board[2][1] == board[2][2])

	vertical := (board[0][0] == board[1][0] && board[1][0] == board[2][0] ||
		board[0][1] == board[1][1] && board[1][01] == board[2][1] ||
		board[0][2] == board[1][2] && board[1][2] == board[2][2])

	diagonal := (board[0][0] == board[1][1] && board[1][1] == board[2][2] ||
		board[2][0] == board[1][1] && board[1][1] == board[0][2])

	return horiz || vertical || diagonal
}
