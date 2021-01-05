package main

import "fmt"

func getAxis(playerName string, axis string) int {
	var x int
	fmt.Printf("%v select a %v: ", playerName, axis)
	fmt.Scanln(&x)

	return x - 1
}
