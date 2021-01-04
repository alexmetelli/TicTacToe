package main

import "fmt"

func getRightSymbol(symbol string, isSymbol bool) string {
	for isSymbol != true {
		fmt.Println("Please select X or O...")
		fmt.Println("Player 1 select your symbol: ")
		fmt.Scanln(symbol)
		isSymbol = isSymbolCorrect(symbol)
	}
	return symbol
}
