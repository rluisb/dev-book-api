package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//Inside init function we generate a random string key to use as a secret for JWT
// func init () {
// 	key := make([]byte, 64)
// 	if _, error := rand.Read(key); error != nil {
// 		log.Fatal(error)
// 	}

// 	secret := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(secret)
// }

func main() {
	config.Load()
	r := router.Generate()

	fmt.Printf("Server running on port: %d \n", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}