package main

import "fmt"

func printBeforeRound(board [3][3]string, round int) {
	printBoard(board)
	fmt.Println()
	fmt.Printf("Round %v\n", round)
}
