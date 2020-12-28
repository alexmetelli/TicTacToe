package main

import "fmt"

// Stores players details
// Name: self descriptive, symbol: 'X' or 'O' at choise.
type player struct {
	name   string
	symbol string
}

func main() {

	//var p1 = player{"Alex", "X"}
	var board = [3][3]string{{"X", "-", "-"},
		{"X", "-", "-"},
		{"-", "-", "-"},
	}
	//var coord = [2]int{1, 2}
	//board = insertToBoard(p1.symbol, board, coord)
	fmt.Println(checkWin(board))
}
