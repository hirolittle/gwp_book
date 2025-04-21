package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeader)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `
<html>
	<header><title>Go Web Programming</title></header>
	<body><h1>hello world</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("Not Implemented"))
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "hiro",
		Threads: []string{"go", "golang"},
	}
	bytes, _ := json.Marshal(post)
	w.Write(bytes)
}
