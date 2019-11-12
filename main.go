package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("starting router")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
