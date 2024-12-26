package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Print("welcome")
	// greeting()
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	fmt.Print("Server running on Localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

// func greeting() {
// 	fmt.Println("hi there")

// }

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>FrontEnd </h1>"))

}
