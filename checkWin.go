package main

// Checks if on of the players won.
// Returns false if no winner true otherwise.
func checkWin(board [3][3]string) bool {
	horiz := checkHorizonatalWin(board)
	vertical := checkVerticalWin(board)
	diagonal := checkDiagonalWin(board)

	return horiz || vertical || diagonal
}
