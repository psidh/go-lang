package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the PizzaCornerTM")
	fmt.Println("Rate our pizza between 1 and 5")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("Thanks for the rating of ", input, "!")
	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Adding 1 to your rating: ", numRating+1)
	}

}
