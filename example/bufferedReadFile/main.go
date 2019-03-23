package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/ahmedalhulaibi/flatfile"
)

// To run this example:
// go run main.go 1
//       or
// go run main.go 2
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

	endOfFile := false

	//To read in the data and marshal using flatfile, there are 2 options
	option, argErr := strconv.ParseInt(os.Args[1], 10, 64)
	checkError(argErr)

	if option == 1 {
		// 1. Append the data read in into a slice until you read the end of the line/record
		//    Then unmarshal into struct
		dataBuffer := []byte("")
		for !endOfFile {
			data, isPrefix, eof := readLine(reader)
			endOfFile = eof
			dataBuffer = append(dataBuffer, data...)
			if !endOfFile {
				if !isPrefix {
					fileRecord := &CustomerRecord{}
					err := flatfile.Unmarshal(dataBuffer, fileRecord, 0, 0)
					fmt.Printf("Unmarshalled: %v\n", fileRecord)
					checkError(err)
					dataBuffer = []byte("")
				}
			}
		}
	} else if option == 2 {
		// 2. incrementally unmarshal using the data we have
		// this is much slower as it requires
		fileRecord := &CustomerRecord{}
		//numFields used to track how many fields were unmarshalled, also used as field offset for next unmarshal
		numFields := 0
		startFieldIndex := 0

		dataBuffer := []byte("")
		remainder := []byte("")
		for !endOfFile {
			data, isPrefix, eof := readLine(reader)
			endOfFile = eof
			if !endOfFile {
				//append read in data to dataBuffer
				dataBuffer = append(dataBuffer, data...)
			}

			//determine how many fields can be unmarshalled
			// store any data from dataBuffer that would not be unmarshalled
			numFields, remainder, err = flatfile.CalcNumFieldsToUnmarshal(dataBuffer, fileRecord, startFieldIndex)
			checkError(err)

			//if we're not at the eof and we can marshal fields using the data in our data buffer
			if !endOfFile && numFields > 0 {
				err := flatfile.Unmarshal(dataBuffer, fileRecord, startFieldIndex, numFields)

				//increment start field index
				startFieldIndex += numFields
				fmt.Printf("Unmarshalled: %v\n", fileRecord)
				checkError(err)

				//store remainder in dataBuffer for future use
				dataBuffer = remainder
			}

			// if we reach EOL
			//reset start field index
			//clear fileRecord
			//clear data buffer
			if !isPrefix {
				startFieldIndex = 0
				fileRecord = &CustomerRecord{}
				dataBuffer = []byte("")
			}
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
