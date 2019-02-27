package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ahmedalhulaibi/ffparser"
)

type CustomerRecord struct {
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

	reader := bufio.NewReader(file)

	eof := false
	for !eof {
		data := readLine(reader)
		if data == nil {
			eof = true
		} else {
			fileHeader := &CustomerRecord{}
			ffparser.Examine(fileHeader)
			err := ffparser.Unmarshal(data, fileHeader, 0)
			fmt.Printf("%v\n", fileHeader)
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
