package main

import (
	"io"
	"net/http"
)

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/dog", logRequest(dogHandler))
	mux.HandleFunc("/cat", logRequest(catHandler))
	mux.HandleFunc("/", logRequest(notFoundHandler))

	return mux
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog")
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
