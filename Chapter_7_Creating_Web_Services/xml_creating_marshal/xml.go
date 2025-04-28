package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      int      `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   int    `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {

	post := Post{
		Id:      1,
		Content: "Hello World",
		Author: Author{
			Id:   1,
			Name: "Hiro",
		},
	}

	output, err := xml.MarshalIndent(post, "", "\t")
	if err != nil {
		fmt.Printf("Error marshalling XML: %v\n", err)
		return
	}

	err = ioutil.WriteFile("Chapter_7_Creating_Web_Services/xml_creating_marshal/post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Printf("Error writing XML file: %v\n", err)
		return
	}

}
