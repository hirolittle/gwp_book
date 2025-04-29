package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/3", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Expected status code 200, got %d", writer.Code)
	}

	var post Post
	err := json.Unmarshal(writer.Body.Bytes(), &post)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %s", err)
	}

	if post.Id != 3 {
		t.Errorf("Expected post ID 3, got %d", post.Id)
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content": "Hello, world!", "author": "tom"}`)
	request, _ := http.NewRequest("PUT", "/post/3", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Expected status code 200, got %d", writer.Code)
	}
}
