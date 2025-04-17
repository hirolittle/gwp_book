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
		Addr:    ":8443",
		Handler: &handler,
	}

	server.ListenAndServeTLS("cert.pem", "key.pem")
}
