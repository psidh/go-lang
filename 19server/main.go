package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// const URL = "http://localhost:3001/get"
// const URL = "http://localhost:3001/post"
const URL = "http://localhost:3001/platform"

func main() {
	// CustomGetRequest(URL)
	// CustomPostRequest(URL)
	// CustomFormData(URL)
	createJSONData()
	// EncodeJson()
}

func CustomGetRequest(URL string) {

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	// fmt.Println(response)

	defer response.Body.Close()

	// fmt.Println("Status Code: ", response.StatusCode)
	// fmt.Println("Content Length: is : ", response.ContentLength)
	var responseStringBuilder strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseStringBuilder.Write(content) // no need store btw
	fmt.Println(byteCount)
	fmt.Println(responseStringBuilder.String())

	//fmt.Println(string(content)) //not industry standard
}

func CustomPostRequest(URL string) {
	requestBody := strings.NewReader(`{
		"hi" : "message_payload",
		"cod" : "modern warfare"
	}`)

	response, err := http.Post(URL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
	defer response.Body.Close()
}

func CustomFormData(u string) {

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")

	response, err := http.PostForm(u, data)
	if err != nil {
		panic(err)
	}
	content, _ := io.ReadAll(response.Body)

	fmt.Print(string(content))

	defer response.Body.Close()

}
