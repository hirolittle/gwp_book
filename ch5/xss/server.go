package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ch5/xss/tmpl.html")
	//t.Execute(w, template.HTML(r.FormValue("comment")))
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ch5/xss/form.html")
	t.Execute(w, nil)
}

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)

	server.ListenAndServe()
}
