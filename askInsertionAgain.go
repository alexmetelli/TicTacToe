package main

import "fmt"

func ankInsertionAgain(p player) [2]int {
	fmt.Printf("%v you can't insert here. Square already taken. Select an empty square\n", p.name)
	newInsertion := takeInput(p)
	return newInsertion
}
