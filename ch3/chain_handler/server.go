package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	helloHandler := HelloHandler{}

	http.Handle("/hello", protect(log(helloHandler)))

	server.ListenAndServe()
}
