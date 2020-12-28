package main

// Stores players details
// Name: self descriptive, symbol: 'X' or 'O' at choise.
type player struct {
	name   string
	symbol string
}

func main() {

	board := makeBoard()
	p1, p2 := askNames()
	game(p1, p2, board)
}
