package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ahmedalhulaibi/ffparser"
)

type CustomerRecord struct {
	//ffparser is one indexed, position starts at 1
	Name        string `ffp:"1,3"`
	OpenDate    string `ffp:"4,10"`
	Age         uint   `ffp:"14,3"`
	Address     string `ffp:"17,15"`
	CountryCode string `ffp:"32,2"`
}

func main() {
	file, err := os.Open("./customers.txt")
	checkError(err)
	defer file.Close()

	reader := bufio.NewReaderSize(file, 1*1)

	endOfFile := false
	dataBuffer := []byte("")

	for !endOfFile {
		data, isPrefix, eof := readLine(reader)
		endOfFile = eof
		dataBuffer = append(dataBuffer, data...)
		if !endOfFile {
			if !isPrefix {
				fileHeader := &CustomerRecord{}
				ffparser.Examine(fileHeader)
				err := ffparser.Unmarshal(dataBuffer, fileHeader, 0)
				fmt.Printf("%v\n", fileHeader)
				checkError(err)
				dataBuffer = []byte("")
			}
			fmt.Println(string(dataBuffer))
		}
	}
	checkError(err)
}

func readLine(reader *bufio.Reader) ([]byte, bool, bool) {
	str, isPrefix, err := reader.ReadLine()
	if err == io.EOF {
		return nil, false, true
	}

	return str, isPrefix, false
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
