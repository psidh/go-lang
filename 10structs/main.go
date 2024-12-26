package main

import "fmt"

type Human struct {
	eyes   int
	legs   int
	name   string // all are private
	UserId int    // public
}

func main() {
	sidharth := Human{2, 2, "Sidharth", 42069}
	fmt.Printf("User %+v\n", sidharth)

	var target int = 5
	arrays := []int{4, 3, 2, 5, 45, 78, 56, 69, 90}

	var left = 0
	var right = len(arrays) - 1

	for left <= right {
		var mid int = left + (right-left)/2

		if arrays[mid] == target {
			fmt.Println("true")
			break
		} else if arrays[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
