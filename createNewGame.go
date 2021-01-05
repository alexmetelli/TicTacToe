package main

func createNewGame() gameObject {
	newGame := gameObject{}
	newGame.state = false
	newGame.round = 0
	newGame.board = makeBoard()

	return newGame
}
