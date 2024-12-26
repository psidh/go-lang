package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices tutorial")

	// declare
	var games = []string{"rdr2", "gtaV", "aoe", "aoc"}
	fmt.Printf("Type : %T \n", games)
	fmt.Println(games)

	games = append(games, "lastofus")

	fmt.Println(games)

	// games = append(games[1:])
	fmt.Println(games[1:3])
	fmt.Println(games[:3])

	highScores := make([]int, 4)

	highScores[0] = 24
	highScores[1] = 945
	highScores[2] = 26
	highScores[3] = 27
	// highScores[4] = 67 panic! can't access and
	//insert on memory that is not yet allocated
	highScores = append(highScores, 555) // valid
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slices using index
	fmt.Println("")
	var courses = []string{"react", "js", "swift","python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
