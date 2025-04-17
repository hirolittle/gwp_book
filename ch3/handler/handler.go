package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, hiro")
}

func main() {

	handler := MyHandler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}

	server.ListenAndServe()

}
