package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("The API is running")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":3000", r))
}