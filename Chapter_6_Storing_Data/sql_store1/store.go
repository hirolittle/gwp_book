package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
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

// get all posts
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, content, author FROM posts limit $1", limit)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Get a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a new post
func (p *Post) Create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.Content, p.Author).Scan(&p.Id)
	return
}

// Update a post
func (p *Post) Update() (err error) {
	_, err = Db.Exec("UPDATE posts SET content = $1, author = $2 WHERE id = $3", p.Content, p.Author, p.Id)
	return
}

// Delete a post
func (p *Post) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM posts WHERE id = $1", p.Id)
	return
}

func main() {

	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	// Create a post
	fmt.Println(post) // {0 Hello World! Sau Sheong}
	post.Create()
	fmt.Println(post) // {1 Hello World! Sau Sheong}

	// Get one post
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost) // {1 Hello World! Sau Sheong}

	// Update the post
	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	// Get all posts
	posts, _ := Posts(10)
	fmt.Println(posts)

	// Delete the post
	readPost.Delete()

	// Get all posts
	posts, _ = Posts(10)
	fmt.Println(posts)
}
