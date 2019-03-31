package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ahmedalhulaibi/flatfile"
)

type CustomerDemographic struct {
	Name     string `flatfile:"1,3"`
	OpenDate string `flatfile:"4,10"`
	Age      uint   `flatfile:"14,3"`
}

type CustomerAddress struct {
	Address     string `flatfile:"1,15"`
	CountryCode string `flatfile:"16,2"`
}

type CustomerRecord struct {
	//flatfile is one indexed, position starts at 1
	Demographics CustomerDemographic `flatfile:"5,16,condition=1-4-CUST"`
	Address      CustomerAddress     `flatfile:"5,17,condition=1-4-ADDR"` //Nested struct. Define the address field range. Within the CustomerAddress struct, ffp tags should start at 1
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
