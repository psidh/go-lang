package main

import "fmt"

func main() {
	fmt.Println("Defer it is huh!")

	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("iHi there")
	myFunc()
}

func myFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
