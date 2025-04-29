package main

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("handle get", func() {

		mux := http.NewServeMux()
		mux.HandleFunc("/post/", handleRequest(&FakePost{}))

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/post/3", nil)
		mux.ServeHTTP(writer, request)
		if writer.Code != 200 {
			GinkgoT().Errorf("Expected status code 200, got %d", writer.Code)
		}

		var post Post
		err := json.Unmarshal(writer.Body.Bytes(), &post)
		if err != nil {
			GinkgoT().Errorf("Error unmarshalling JSON: %s", err)
		}

		if post.Id != 3 {
			GinkgoT().Errorf("Expected post ID 3, got %d", post.Id)
		}
	})
})
