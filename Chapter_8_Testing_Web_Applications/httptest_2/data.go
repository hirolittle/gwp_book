package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp password=123456 dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Get a post
func Retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a post
func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update a post
func (post *Post) Update() (err error) {
	_, err = Db.Exec("UPDATE posts SET content = $1, author = $2 WHERE id = $3", post.Content, post.Author, post.Id)
	return
}

// Delete a post
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM posts WHERE id = $1", post.Id)
	return
}
