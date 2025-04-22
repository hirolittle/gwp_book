package main

import (
	"html/template"
	"math/rand"
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

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("ch5/random_number/tmpl.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)

}
