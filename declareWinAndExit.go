package main

import (
	"fmt"
	"os"
)

func declareWinAndExit(pName string) {
	fmt.Printf("%v WON!\n", pName)
	os.Exit(0)
}
