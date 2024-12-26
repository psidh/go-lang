package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// createJSONData()
	consumeJSON()
}

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func createJSONData() {
	courses := []course{
		{
			Name:     "react",
			Price:    200,
			Platform: "udemy",
			Password: "123",
			Tags:     []string{"web", "front-end"},
		},
		{
			Name:     "angular",
			Price:    250,
			Platform: "udemy",
			Password: "456",
			Tags:     []string{"google", "front-end"},
		},
		{
			Name:     "go-lang",
			Price:    150,
			Platform: "udemy",
			Password: "789",
			Tags:     nil,
		},
	}

	// package this as jSON

	finalJSON, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)

}

func consumeJSON() {
	jsonSampleData := []byte(`
	{
			"Name":     "react",
			"Price":    200,
			"Platform": "udemy",
			"Password": "123",
			"Tags": ["web", "front-end"]
	}
	`)

	var storeJsonCourse course
	checkValid := json.Valid(jsonSampleData)
	if checkValid {
		json.Unmarshal(jsonSampleData, &storeJsonCourse)
		// fmt.Printf("%#v\n", storeJsonCourse)
	} else {
		fmt.Print("Error")
	}

	// some cases where we need to add data to key value pair

	var myOnlineDataMap map[string]interface{}

	json.Unmarshal(jsonSampleData, &myOnlineDataMap)
	fmt.Printf("%#v\n", myOnlineDataMap)
	for k, v := range myOnlineDataMap {
		fmt.Printf("Key is %v and value is %v", k, v)
		fmt.Println("")
	}
}
