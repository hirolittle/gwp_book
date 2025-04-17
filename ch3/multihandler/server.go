package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

type GoodbyeHandler struct{}

func (h *GoodbyeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye")
}

func main() {

	helloHandler := HelloHandler{}
	goodbyeHandler := GoodbyeHandler{}

	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/hello", &helloHandler)
	http.Handle("/goodbye", &goodbyeHandler)

	server.ListenAndServe()
}
