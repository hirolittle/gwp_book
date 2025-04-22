package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("ch5/trigger_template/tmpl.html")
	if err != nil {
		fmt.Printf("模板解析错误: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, "hello world")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
