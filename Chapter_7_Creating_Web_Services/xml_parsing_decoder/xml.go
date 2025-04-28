package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       int       `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   int    `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      int    `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("Chapter_7_Creating_Web_Services/xml_parsing_unmarshal_2/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "post" {
				var post Post
				decoder.DecodeElement(&post, &se)
				fmt.Printf("%#v\n", post)
			}
		}
	}
}
