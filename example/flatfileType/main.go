package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ahmedalhulaibi/flatfile"
)

type CustomerRecord struct {
	//flatfile is one indexed, position starts at 1
	Name        string `flatfile:"1,3"`
	OpenDate    string `flatfile:"4,10"`
	Age         uint   `flatfile:"14,3"`
	Address     string `flatfile:"17,15"`
	CountryCode string `flatfile:"32,2"`
}

func main() {
	file, err := os.Open("./customers.txt")
	checkError(err)
	defer file.Close()

	//Creare buffered reader with small buffer size to simulate reading long lines of data that exceed a buffer limit
	reader := bufio.NewReaderSize(file, 1*1)

	fileRecord := &CustomerRecord{}
	customerFile, err := flatfile.New(reader, fileRecord)
	for err != io.EOF {
		err = customerFile.Read()
		fmt.Printf("Unmarshalled: %v\n", fileRecord)
		checkError(err)
	}

}

func checkError(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}
