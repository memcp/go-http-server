package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Listening on 127.0.0.1:8080")

	routes := map[string]func(w http.ResponseWriter, r *http.Request){
		"/dog": func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s\n", r.Method, r.URL.Path)
			io.WriteString(w, "dog")
		},
		"/cat": func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s\n", r.Method, r.URL.Path)
			io.WriteString(w, "cat")
		},
	}

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		if _, ok := routes[r.URL.Path]; !ok {
			log.Printf("%s %s page not found\n", r.Method, r.URL.Path)
			http.NotFound(w, r)
			return
		}
	}
	http.HandleFunc("/", indexHandler)

	for route := range routes {
		http.HandleFunc(route, routes[route])
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
