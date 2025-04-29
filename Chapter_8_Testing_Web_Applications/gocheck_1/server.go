package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

func main() {

	db, err := sql.Open("postgres", "user=gwp password=123456 dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))

	server.ListenAndServe()

}

// main handle request
func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)

		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Create a post
// Post /post/
func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	err = json.Unmarshal(body, &post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = post.create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	return
}

// Get a post
// Get /post/:id
func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = post.fetch(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// Update a post
// Put /post/:id
func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, fmt.Sprintf("id: %d, err: %s", id, err.Error()), http.StatusInternalServerError)
		return
	}

	err = post.fetch(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("post: %v, err: %s", post, err.Error()), http.StatusInternalServerError)
		return
	}
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	err = json.Unmarshal(body, &post)
	if err != nil {
		http.Error(w, fmt.Sprintf("unmarshal failed, post: %v, err: %s", post, err.Error()), http.StatusInternalServerError)
		return
	}
	err = post.update()
	if err != nil {
		http.Error(w, fmt.Sprintf("update failed, err: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

// Delete a post
// Delete /post/:id
func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = post.fetch(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
