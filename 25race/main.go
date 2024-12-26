package main

import (
	"fmt"
	"sync"
)

func main() {

	var score = []int{0}
	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("One")
		mutex.Lock()
		score =
			append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("Two")
		mutex.Lock()
		score =
			append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("Three")
		mutex.Lock()
		score =
			append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("Three")
		mutex.RLock()
		fmt.Println(score)
		mutex.RUnlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(score)
}
