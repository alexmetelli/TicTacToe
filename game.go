package main

import "fmt"

func game(p1 player, p2 player, board [3][3]string) {

	status := checkWin(board)
	round := 0

	for status != true {

		round++
		printBoard(board)
		fmt.Println()

		fmt.Printf("Round %v\n", round)
		pOneInput := takeInput(p1)

		// Check if input doesn't overlap exixting entries
		valid := checkInsertion(pOneInput, board)
		fmt.Println(valid)
		for valid == false {
			fmt.Printf("%v you can't insert here. Square already taken. Select an empty square\n", p1.name)
			pOneInput = takeInput(p1)
			valid = checkInsertion(pOneInput, board)
		}

		board = insertToBoard(p1.symbol, board, pOneInput)
		fmt.Println()
		printBoard(board)

		status = checkWin(board)
		if status == true {
			fmt.Printf("%v Wins!!!\n", p1.name)
			break
		} else {
			pTwoInput := takeInput(p2)
			valid = checkInsertion(pTwoInput, board)
			fmt.Println(valid)
			for valid == false {
				fmt.Printf("%v you can't insert here. Square already taken. Select an empty square\n", p2.name)
				pTwoInput = takeInput(p2)
				valid = checkInsertion(pTwoInput, board)
			}

			board = insertToBoard(p2.symbol, board, pTwoInput)
			fmt.Println()

			status = checkWin(board)
			if status == true {
				fmt.Printf("%v Wins!!!\n", p2.name)
				break
			}
		}

	}
}
