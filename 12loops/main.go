package main

import "fmt"

func main() {
	fmt.Println("Loops")

//	days := []string{"S", "M", "T", "W", "TH", "FR", "ST"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	
	// for index, day := range days {
	// 	fmt.Println("Index is : %v and value is %v", index, day)
	// }
	rogue := 1
	for rogue < 10{
		fmt.Println(rogue)
		rogue++
	}
}
