package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")
	one := 2
	//declaration of pointer
	var ptr *int
	//declaration of pointer and getting reference of a memory address
	var ptr2 *int = &one
	fmt.Println(ptr)
	fmt.Println(ptr2)
	// dereferencing the pointer and getting the value
	fmt.Println(*ptr2 * 8)
}
