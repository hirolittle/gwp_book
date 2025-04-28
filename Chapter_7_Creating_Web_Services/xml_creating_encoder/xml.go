package main

import (
	"encoding/xml"
	"fmt"
	"os"
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

	xmlFile, err := os.Create("Chapter_7_Creating_Web_Services/xml_creating_encoder/post.xml")
	if err != nil {
		fmt.Printf("Error creating XML file: %v\n", err)
		return
	}
	defer xmlFile.Close()
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Printf("Error encoding XML: %v\n", err)
		return
	}
}
