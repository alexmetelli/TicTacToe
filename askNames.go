package main

import (
	"fmt"
	"strings"
)

// Asks both players to input their names and to player 1 to choose
// a symbol betweek 'X' or 'O'.
// Return p1 and p2 as player type.
func askNames() (p1 player, p2 player) {

	fmt.Println("Enter player's 1 name: ")
	fmt.Scanln(&p1.name)
	fmt.Println("Player 1 select your symbol: ")
	fmt.Scanln(&p1.symbol)
	p1.symbol = strings.ToUpper(p1.symbol)
	symbolCheck := p1.symbol == "X" || p1.symbol == "O"

	for symbolCheck != true {
		fmt.Println("Player 1 select your symbol: ")
		fmt.Scanln(&p1.symbol)
		symbolCheck = p1.symbol == "X" || p1.symbol == "O"
	}

	fmt.Println("Enter player's 2 name: ")
	fmt.Scanln(&p2.name)
	if p1.symbol == "X" {
		p2.symbol = "O"
	} else {
		p2.symbol = "X"
	}

	return p1, p2

}
