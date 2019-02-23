package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ahmedalhulaibi/ffparser"
)

type FileHeader struct {
	LogicalRecordTypeID string `ffp:"1,1"`
	LogicalRecordCount  uint32 `ffp:"2,9"`
	NameIdentifier      string `ffp:"11,10"`
}

func main() {
	file, err := os.Open("./test.txt")
	checkError(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	eof := false
	for !eof {
		data := readLine(reader)
		if data == nil {
			eof = true
		} else {
			fileHeader := &FileHeader{}
			ffparser.Examine(fileHeader)
			err := ffparser.Unmarshal(data, fileHeader, 0)
			fmt.Printf("%#v\n", fileHeader)
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
