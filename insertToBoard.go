package main

// symbol: string,  player symbol to be found in player.symbol.
// board: 2D array, current state of the board.
// coord: array, coordinates for insertion chosen by the player.
// Returns 2D array with updated state of the board
func insertToBoard(symbol string, board [3][3]string, coords [2]int) [3][3]string {

	y := coords[1]
	x := coords[0]
	board[x][y] = symbol

	return board
}
