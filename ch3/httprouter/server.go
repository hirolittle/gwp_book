package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s", ps.ByName("name"))
}

func main() {

	mux := httprouter.New()

	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
