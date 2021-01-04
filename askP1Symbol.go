package main

import (
	"fmt"
	"strings"
)

func askP1Symbol(p player) string {
	fmt.Println("Player 1 select your symbol: ")
	fmt.Scanln(&p.symbol)
	p.symbol = strings.ToUpper(p.symbol)

	return p.symbol
}
