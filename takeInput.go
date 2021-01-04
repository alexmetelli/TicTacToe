package main

import (
	"fmt"
	"os"
)

// Asks the user to enter coordinates for insertion.
// Return an array with coordinates in the form: [column, row]
func takeInput(plr player) [2]int {

	var coordinates [2]int
	var exitGame string
	// Takes coordinate for x axis
	fmt.Printf("%v select a row: ", plr.name)
	//fmt.Scanln(&exitGame)
	fmt.Scanln(&coordinates[0], &exitGame)
	if exitGame == "Q" {
		fmt.Println("Quitting game...")
		os.Exit(1)
	} else {
		coordinates[0] = coordinates[0] - 1
	}

	cond1 := coordinates[0] < 0 || coordinates[0] > 2
	// checks if coordinates entered are between bounds (1,3)
	for cond1 == true {
		fmt.Println("Error! Coordinate must be between 1 and 3.")
		fmt.Printf("%v select a row: ", plr.name)
		fmt.Scanln(&coordinates[0])
		coordinates[0] = coordinates[0] - 1
		cond1 = coordinates[0] < 0 || coordinates[0] > 2
	}
	// Takes coordinate for y axis
	fmt.Printf("%v select a column: ", plr.name)
	fmt.Scanln(&coordinates[1])
	coordinates[1] = coordinates[1] - 1
	cond2 := coordinates[1] < 0 || coordinates[1] > 2
	// checks if coordinates entered are between bounds (1,3)
	for cond2 == true {
		fmt.Println("Error! Coordinate must be between 1 and 3.")
		fmt.Printf("%v select a column: ", plr.name)
		fmt.Scanln(&coordinates[1])
		coordinates[1] = coordinates[1] - 1
		cond2 = coordinates[1] < 0 || coordinates[1] > 2
	}

	return coordinates
}
