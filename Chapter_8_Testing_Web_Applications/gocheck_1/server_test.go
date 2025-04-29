package main

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PostTestSuite struct{}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuite) TestHandleGet(c *C) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/post/3", nil)
	mux.ServeHTTP(writer, request)
	c.Assert(writer.Code, Equals, 200)

	var post Post
	err := json.Unmarshal(writer.Body.Bytes(), &post)
	c.Assert(err, IsNil)
	c.Assert(post.Id, Equals, 3)
}
