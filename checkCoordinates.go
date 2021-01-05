package main

import "fmt"

func checkCoordinates(coord int, playerName, axis string) int {
	cond1 := coord < 0 || coord > 2

	for cond1 == true {
		fmt.Println("Error! Coordinate must be between 1 and 3.")
		coord = getAxis(playerName, axis)
		cond1 = coord < 0 || coord > 2
	}

	return coord
}
