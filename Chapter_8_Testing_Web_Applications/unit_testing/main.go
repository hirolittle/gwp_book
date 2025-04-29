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

// decode json from file to struct
func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}
	return
}

func unmarshal(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
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

	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}
	return
}

func main() {
	post, err := decode("Chapter_8_Testing_Web_Applications/unit_testing/post.json")
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", post)

	post, err = unmarshal("Chapter_8_Testing_Web_Applications/unit_testing/post.json")
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
	}
	fmt.Printf("%#v\n", post)
}
