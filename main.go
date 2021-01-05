package main

// Stores players details
// Name: self descriptive, symbol: 'X' or 'O' at choise.
type player struct {
	name   string
	symbol string
}

type gameObject struct {
	board [3][3]string
	state bool
	round int
	p1    player
	p2    player
}

func main() {

	newGame := createNewGame()
	newGame.p1, newGame.p2 = askPlayersDetails()
	for newGame.state == false {
		newGame.state, newGame.board = round(newGame)
	}
}
