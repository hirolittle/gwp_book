package main

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finished test - ", c.TestName())
}

func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Post test suite")
}

func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finishing Post test suite")
}

func (s *PostTestSuite) TestHandleGet(c *C) {

	request, _ := http.NewRequest("GET", "/post/3", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)

	var post Post
	err := json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Assert(err, IsNil)
	c.Assert(post.Id, Equals, 3)
}

func (s *PostTestSuite) TestHandlePut(c *C) {
	jsonData := strings.NewReader(`{"content": "Hello, world!", "author": "tom"}`)
	request, _ := http.NewRequest("PUT", "/post/3", jsonData)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)

	c.Check(s.post.Content, Equals, "Hello, world!")
	c.Check(s.post.Author, Equals, "tom")

}
