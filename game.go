package main

import "fmt"

func game(p1 player, p2 player, board [3][3]string) {

	status := checkWin(board)

	for status != true {

		printBoard(board)
		fmt.Println("//\\//\\//\\//\\///\\//\\")
		pOneInput := takeInput(p1)
		board = insertToBoard(p1.symbol, board, pOneInput)
		printBoard(board)
		fmt.Println("//\\//\\//\\//\\///\\//\\")
		status = checkWin(board)
		if status == true {
			fmt.Printf("%v Wins!!!\n", p1.name)
			break
		} else {
			pTwoInput := takeInput(p2)
			board = insertToBoard(p2.symbol, board, pTwoInput)
			printBoard(board)
			fmt.Println("//\\//\\//\\//\\///\\//\\")
			status = checkWin(board)
			if status == true {
				fmt.Printf("%v Wins!!!\n", p2.name)
				break
			}
		}
	}
}
