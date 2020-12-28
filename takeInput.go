package main

import (
	"fmt"
)

// Asks the user to enter coordinates for insertion.
// Return an array with coordinates in the form: [column, row]
func takeInput(plr player) [2]int {

	var coordinates [2]int
	fmt.Println("%v select a column: ", plr.name)
	fmt.Scanln(&coordinates[0])
	fmt.Println("%v select a row: ", plr.name)
	fmt.Scanln(&coordinates[1])

	return coordinates
}
