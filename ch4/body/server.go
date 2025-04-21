package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/body", body)

	server.ListenAndServe()

}

func body(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	bodyBytes := make([]byte, length)
	r.Body.Read(bodyBytes)
	bodyString := string(bodyBytes)
	fmt.Fprintf(w, "Body: %v\n", bodyString)
}
