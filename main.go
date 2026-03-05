package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on 127.0.0.1:8080")

	mux := setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
