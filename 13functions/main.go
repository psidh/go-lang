package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search() {
	fmt.Println("Enter a number : ")
}
func factorial(num int) int {
	if num == 1 {
		return num
	}
	return num * factorial(num-1)

}
func proAdder(values ...int) int {

	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func main() {
	fmt.Println("Main Function")
	search()
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	var ans int = factorial(num)
	fmt.Printf("Factorial of %v is %v", num, ans)
	fmt.Println("")
	var ans2 int = proAdder(6, 7, 8, 10)
	fmt.Println(ans2)

}
