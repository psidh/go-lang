package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to GITAM"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Indi Tandoori Paneer!:")

	// comma ok syntax for the error handling
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the ratings!", input)
	fmt.Printf("Thanks for the ratings! %T", input)
}
