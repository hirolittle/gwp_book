package main

import (
	"html/template"
	"net/http"
	"time"
)

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()

}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {

	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("ch5/custom_function/tmpl.html")
	t.Execute(w, time.Now())

}
