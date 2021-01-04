package main

func setP2Symbol(p1symbol string) string {
	var p2symbol string
	if p1symbol == "X" {
		p2symbol = "O"
	} else {
		p2symbol = "X"
	}
	return p2symbol
}
