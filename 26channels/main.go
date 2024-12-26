package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels 1.0")
	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	//Receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, IsChannelOpen := <-myChannel

		fmt.Println(IsChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myChannel)
		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	//Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(myChannel)
		// myChannel <- 0
		// myChannel <- 6
		wg.Done()
	}(myChannel, wg)

	wg.Wait()

	// myChannel <- 5
	// fmt.Println(<-myChannel)

}
