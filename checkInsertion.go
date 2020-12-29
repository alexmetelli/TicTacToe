package main

// Check if "square" at insertion coordinates is empty.
// coord: coordionates to check (y, x), board: current board status
// return true if insertion space is empty false otherwise
func checkInsertion(coord [2]int, board [3][3]string) bool {

	y := coord[1]
	x := coord[0]
	cond := board[x][y] != "X" && board[x][y] != "O"

	return cond

}
