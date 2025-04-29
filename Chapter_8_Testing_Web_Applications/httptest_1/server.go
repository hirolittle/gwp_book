package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/post/", handleRequest)

	server.ListenAndServe()

}

// main handle request
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Create a post
// Post /post/
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = post.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Post created")
	w.WriteHeader(http.StatusOK)
	return
}

// Get a post
// Get /post/:id
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post, err := Retrieve(id)
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
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, fmt.Sprintf("id: %d, err: %s", id, err.Error()), http.StatusInternalServerError)
		return
	}

	post, err := Retrieve(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("post: %v, err: %s", post, err.Error()), http.StatusInternalServerError)
		return
	}
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	post.Id = id
	err = json.Unmarshal(body, &post)
	if err != nil {
		http.Error(w, fmt.Sprintf("unmarshal failed, post: %v, err: %s", post, err.Error()), http.StatusInternalServerError)
		return
	}
	err = post.Update()
	if err != nil {
		http.Error(w, fmt.Sprintf("update failed, err: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

// Delete a post
// Delete /post/:id
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := Retrieve(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
