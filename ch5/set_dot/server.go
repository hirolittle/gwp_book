package main

import (
	"html/template"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()

}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ch5/set_dot/tmpl.html")
	t.Execute(w, "hello")
}
