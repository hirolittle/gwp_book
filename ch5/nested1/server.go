package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ch5/nested1/layout.html")
	t.ExecuteTemplate(w, "layout", "")
}

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()

}
