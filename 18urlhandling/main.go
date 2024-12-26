package main

import (
	"fmt"
	"strings"

	// "net/http"
	"net/url"
)

const URL = "https://fuck-three.vercel.app/learn?coursename=reactjs&paymentId=123"

func main() {
	fmt.Println("URL Handling")

	// response, err := http.Get(URL)
	// catch(err)

	//let us parse the using URL
	result, err := url.Parse(URL)
	catch(err)
	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(strings.Split(result.RawQuery, "&")[0])

	queryParams := result.Query()

	fmt.Printf("Type of query params are: %T", queryParams)
	fmt.Println("Type of query params are: ", queryParams["coursename"])

	for _, val := range queryParams {
		fmt.Println("Param is :", val)
	}

	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=hitesh",
	}

	creatingTheURL := partsOfURL.String()
	fmt.Print(creatingTheURL)
}

func catch(e error) {
	if e != nil {
		panic(e)
	}
}
