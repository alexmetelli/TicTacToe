package main

// Asks both players to input their names and to player 1 to choose
// a symbol betweek 'X' or 'O'.
// Return p1 and p2 as player type.
func askPlayersDetails(p1, p2 player) (player, player) {

	p1.name = askPlayerName(p1)
	p1.symbol = askP1Symbol(p1)
	symbolCheck := isSymbolCorrect(p1.symbol)
	if symbolCheck == false {
		p1.symbol = getRightSymbol(p1.symbol, symbolCheck)
	}

	p2.name = askPlayerName(p2)
	p2.symbol = setP2Symbol(p1.symbol)
	return p1, p2

}
