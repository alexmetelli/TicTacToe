package main

import (
	"fmt"
)

func printBoard(board [3][3]string) {

	for i := 0; i < 3; i++ {
		fmt.Println("%v\t%v\t%v\t", board[i][0], board[i][1], board[i][2])
	}
}
