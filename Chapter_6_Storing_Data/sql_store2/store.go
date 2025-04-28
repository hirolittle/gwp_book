package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp password=123456 dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (comment Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not fount")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		return
	}
	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	return

}

// Create a post
func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	post.Create()

	// Add a comment
	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)                  // {1 Hello World! Sau Sheong [{1 Good post! Joe 0x140000ba000}]}
	fmt.Println(readPost.Comments)         // [{1 Good post! Joe 0x140000ba000}]
	fmt.Println(readPost.Comments[0].Post) // &{1 Hello World! Sau Sheong [{1 Good post! Joe 0x140000ba000}]}
}
