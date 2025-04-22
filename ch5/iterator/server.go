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
	t, _ := template.ParseFiles("ch5/iterator/tmpl.html")
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	t.Execute(w, daysOfWeek)
}
