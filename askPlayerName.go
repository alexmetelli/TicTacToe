package main

import (
	"fmt"
)

func askPlayerName(p player) string {
	fmt.Printf("Enter player's %v name: ", p)
	fmt.Scanln(&p.name)

	return p.name
}
