package main

// Creates a 3X3 "empty" board with dashes only.
// Return 2D string array
func makeBoard() [3][3]string {

	board := [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	return board
}
