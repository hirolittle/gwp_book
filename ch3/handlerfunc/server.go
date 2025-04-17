package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
