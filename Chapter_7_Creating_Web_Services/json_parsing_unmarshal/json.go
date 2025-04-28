package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	jsonFile, err := os.Open("Chapter_7_Creating_Web_Services/json_parsing_unmarshal/post.json")
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}

	var post Post
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	fmt.Printf("%#v\n", post)
}
