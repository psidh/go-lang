package main

import (
	"fmt"
	"log"
	"mongodb-backend/router"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at Port:4000")
}
