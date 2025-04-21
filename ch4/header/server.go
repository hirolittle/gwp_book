package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/headers", headers)

	server.ListenAndServe()

}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintf(w, "Headers: %v\n", h)

	a := h["Accept-Encoding"]
	fmt.Fprintf(w, "Accept-Encoding: %v\n", a)

	b := h.Get("Accept-Encoding")
	fmt.Fprintf(w, "Accept-Encoding: %v\n", b)
}
