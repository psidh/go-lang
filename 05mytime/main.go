package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hi there at this moment to study golang!")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.May, 12, 23, 0, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("02-01-2006"))

}
