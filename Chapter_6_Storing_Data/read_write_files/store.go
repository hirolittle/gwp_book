package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	data := []byte("hello world")

	// write to file and read from file using WriteFile and ReadFile
	err := ioutil.WriteFile("Chapter_6_Storing_Data/read_write_files/data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile("Chapter_6_Storing_Data/read_write_files/data1")
	fmt.Println(string(read1))

	// write to file and read from file using the File struct
	file1, _ := os.Create("Chapter_6_Storing_Data/read_write_files/data2")
	defer file1.Close()

	bytes1, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes1)

	file2, _ := os.Open("Chapter_6_Storing_Data/read_write_files/data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	byte2, _ := file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", byte2)
	fmt.Println(string(read2))

}
