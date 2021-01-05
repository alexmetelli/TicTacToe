package main

import (
	"fmt"
)

func askPlayerName(p player) string {
	fmt.Println("Enter player's name: ")
	fmt.Scanln(&p.name)

	return p.name
}
