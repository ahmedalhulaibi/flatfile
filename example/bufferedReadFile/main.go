package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/ahmedalhulaibi/ffparser"
)

// To run this example:
// go run main.go 1
//       or
// go run main.go 2
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

	//Creare buffered reader with small buffer size to simulate reading long lines of data that exceed a buffer limit
	reader := bufio.NewReaderSize(file, 1*1)

	endOfFile := false

	//To read in the data and marshal using ffparser, there are 2 options
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
					//ffparser.Examine(fileRecord)
					//fmt.Println("Before Unmarshal Buffer: ", string(dataBuffer))
					err := ffparser.Unmarshal(dataBuffer, fileRecord, 0, 0)
					//fmt.Printf("Unmarshalled: %v\n", fileRecord)
					checkError(err)
					dataBuffer = []byte("")
				}
				//fmt.Println("After Unmarshal Buffer: ", string(dataBuffer))
			}
		}
	} else if option == 2 {
		// 2. incrementally unmarshal using the data we have
		// this is much slower as it makes many repeated calls
		fileRecord := &CustomerRecord{}
		//used to track how many fields were unmarshalled, also used as field offset for next unmarshal
		numFields := 0
		startFieldIndex := 0
		dataBuffer := []byte("")
		remainder := []byte("")
		for !endOfFile {
			data, isPrefix, eof := readLine(reader)
			endOfFile = eof
			if !endOfFile {
				dataBuffer = append(dataBuffer, data...)
			}

			numFields, remainder, err = ffparser.CalcNumFieldsToMarshal(dataBuffer, fileRecord, startFieldIndex)
			//fmt.Printf("Numfields(%s,%v,%d) = %d, %s\n", string(dataBuffer), fileRecord, startFieldIndex, numFields, string(remainder))
			checkError(err)
			//if we're not at the eof and we can marshal fields using the data in our data buffer
			if !endOfFile && numFields > 0 {
				//ffparser.Examine(fileRecord)
				//fmt.Printf("Unmarshal(%s,%v,%d)\n", string(dataBuffer), fileRecord, startFieldIndex)
				//fmt.Println("Before Unmarshal Buffer: ", string(dataBuffer))
				err := ffparser.Unmarshal(dataBuffer, fileRecord, startFieldIndex, numFields)
				startFieldIndex += numFields
				fmt.Printf("Unmarshalled: %v\n", fileRecord)
				checkError(err)
				dataBuffer = remainder
				//fmt.Println("After Unmarshal Buffer: ", string(dataBuffer))
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
