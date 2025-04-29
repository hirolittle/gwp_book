package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
}

func teardown() {}

func TestHandleGet(t *testing.T) {
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
	jsonData := strings.NewReader(`{"content": "Hello, jim!", "author": "jim"}`)
	request, _ := http.NewRequest("PUT", "/post/3", jsonData)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Expected status code 200, got %d", writer.Code)
	}
}
