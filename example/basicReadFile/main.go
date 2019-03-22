package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ahmedalhulaibi/ffparser"
)

// To run this example:
// go run main.go

type CustomerRecord struct {
	//ffparser is one indexed, position starts at 1
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

	reader := bufio.NewReader(file)

	eof := false
	for !eof {
		data := readLine(reader)
		if data == nil {
			eof = true
		} else {
			fileRecord := &CustomerRecord{}
			//unmarhsal text data to struct
			err := ffparser.Unmarshal(data, fileRecord, 0, 0)
			fmt.Printf("%v\n", fileRecord)
			checkError(err)
		}
	}
	checkError(err)
}

func readLine(reader *bufio.Reader) []byte {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return nil
	}

	return str
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
