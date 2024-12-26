package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals = []string{}
var mutex sync.Mutex

func main() {
	websiteList := []string{
		"https://vercel.com",
		"https://google.com",
		"https://go.dev",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	response, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		fmt.Printf("%d, status code for %s\n", response.StatusCode, endpoint)
	}

}
