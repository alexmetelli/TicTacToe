package main

func round(game gameObject) (bool, [3][3]string) {

	game.round++
	// Player one turn.
	pOneInput := takeInput(game.p1)
	valid := checkInsertion(pOneInput, game.board)
	for valid == false {
		pOneInput = ankInsertionAgain(game.p1)
		valid = checkInsertion(pOneInput, game.board)
	}
	game.board = insertToBoard(game.p1.symbol, game.board, pOneInput)
	printBoard(game.board)
	game.state = checkWin(game.board)
	if game.state {
		declareWinAndExit(game.p1.name)
	}

	// Player two turn.
	pTwoInput := takeInput(game.p2)
	valid = checkInsertion(pTwoInput, game.board)
	for valid == false {
		pTwoInput = ankInsertionAgain(game.p2)
		valid = checkInsertion(pTwoInput, game.board)
	}
	game.board = insertToBoard(game.p2.symbol, game.board, pTwoInput)
	printBoard(game.board)
	game.state = checkWin(game.board)
	if game.state {
		declareWinAndExit(game.p2.name)
	}

	return game.state, game.board

}
