package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

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
