package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://fuck-three.vercel.app/"

func main() {
	fmt.Println("Web Requests")
	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type : %T\n", response)

	defer response.Body.Close() // Should close it
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Print(string(body))
	}
}
