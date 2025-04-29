package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type Post struct {
	Db      *sql.DB
	Id      int
	Content string
	Author  string
}

// Get a post
func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("SELECT id, content, author FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a post
func (post *Post) create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"
	stmt, err := post.Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = post.Db.Exec("UPDATE posts SET content = $1, author = $2 WHERE id = $3", post.Content, post.Author, post.Id)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = post.Db.Exec("DELETE FROM posts WHERE id = $1", post.Id)
	return
}
