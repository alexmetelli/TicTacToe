package main

// Asks the user to enter coordinates for insertion.
// Return an array with coordinates in the form: [column, row]
func takeInput(p player) [2]int {

	var coordinates [2]int
	// Takes X axis coordinates and checks for validity.
	coordinates[0] = getAxis(p.name, "row")
	coordinates[0] = checkCoordinates(coordinates[0], p.name, "row")
	// Takes Y axis coordinates and checks for validity.
	coordinates[1] = getAxis(p.name, "column")
	coordinates[1] = checkCoordinates(coordinates[1], p.name, "column")

	return coordinates
}
