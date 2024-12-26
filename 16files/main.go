package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to handling files in GOlang")
	content := "This needs to go in the file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	} else {
		fmt.Print("This is the length btw: ", length)
	}
	readFile("./myfile.txt")
	defer file.Close() // recommended use of defer

}

func readFile(f string) {
	file, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Text Data inside is \n", string(file))
	}
}

func checkNilError(e error) {
	if e != nil {
		panic(e)
	}
}
