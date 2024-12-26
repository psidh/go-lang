package main

import "fmt"

type Human struct {
	eyes   int
	legs   int
	name   string // all are private
	UserId int    // public
}

func (h Human) getGreeting() {
	fmt.Printf("H has %v eyes", h.eyes)
	fmt.Printf("H has %v legs", h.legs)
}

func (h *Human) newId() { // a pointer references the struct created and
	// hence manipulates the values else it just gets a copy of the value
	h.UserId = 45

}

func main() {
	sidharth := Human{2, 2, "Sidharth", 42069}
	// fmt.Printf("User %+v\n", sidharth)
	// sidharth.getGreeting()
	// fmt.Println("")
	sidharth.UserId = 90
	fmt.Printf("User Id before : %v", sidharth.UserId)
	fmt.Println("")
	sidharth.newId()
	fmt.Println("")
	fmt.Printf("User Id after : %v", sidharth.UserId)
}
