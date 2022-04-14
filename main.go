package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	key := make([]byte, 64)
// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	keyBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(keyBase64)
// }

func main() {
	fmt.Println("Rodando API")

	config.Load()

	log.Fatal(http.ListenAndServe(":5500", router.NewRouter()))
}
