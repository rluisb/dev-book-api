package main

import (
	"fmt"
	"ip-search-app/app"
	"log"
	"os"
)

func main () {
	fmt.Println("Ignition point")

	application := app.Generate()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}


}