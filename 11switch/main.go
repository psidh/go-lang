package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Game in golang")

	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
	case 6:
		fmt.Println("Open the die")
	case 1:
		fmt.Println("Move 1")
	case 2:
		fmt.Println("Move 2")
	case 3:
		fmt.Println("Move 3")
	case 4:
		fmt.Println("Move 4")
	case 5:
		fmt.Println("Move 5")
	default:
		fmt.Println("Error")
	}
}
