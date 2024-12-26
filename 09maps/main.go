package main

import "fmt"

func main() {
	var mymap = make(map[string]string)

	mymap["js"] = "JavaScript"
	mymap["ts"] = "TypeScript"
	mymap["cs"] = "CoffeeScript"
	fmt.Println(mymap)
	
	delete(mymap ,"js")
	fmt.Println(mymap)
}
