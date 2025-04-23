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
	t, _ := template.ParseFiles("ch5/context_aware/tmpl.html")
	content := `I asked <i>"What's up?""</i>`
	t.Execute(w, content)
}
