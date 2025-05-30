package main

import (
	"fmt"
	"io"
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
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}
