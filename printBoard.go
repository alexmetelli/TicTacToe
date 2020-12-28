package main

import (
	"fmt"
)

// Prints the board to square format.
func printBoard(board [3][3]string) {

	for i := 0; i < 3; i++ {
		fmt.Printf("%v\t%v\t%v\t\n", board[i][0], board[i][1], board[i][2])
		fmt.Println()
	}
}
