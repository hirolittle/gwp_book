package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {

	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
				Id:      1,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			Comment{
				Id:      2,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	jsonFile, err := os.Create("Chapter_7_Creating_Web_Services/json_creating_encoder/post.json")
	if err != nil {
		fmt.Printf("Error creating JSON file: %v\n", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
	}
}
