package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	var post Post
	err = xml.Unmarshal(xmlData, &post)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}
	fmt.Printf("%#v\n", post)
}
