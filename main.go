package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Listening on 127.0.0.1")

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("INFO: ", "GET /")
		io.WriteString(w, "hello, go!")
	}

	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
