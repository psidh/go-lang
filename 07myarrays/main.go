package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays section")

	var fruits [10]string

	fruits[0] = "Apple"
	fruits[1] = "PineApple"
	fruits[4] = "Banana"
	fmt.Println(fruits)
	fmt.Println(len(fruits)) // gives 10

	veg := [3]string{"potato", "mushroom", "olives"}

	fmt.Println(veg)
	fmt.Println(len(veg))

}
