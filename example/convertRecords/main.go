package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/ahmedalhulaibi/ffparser"
)

// To run this example:
// go run main.go
type CustomerDemographic struct {
	Name     string `flatfile:"1,3" json:"name" xml:"name"`
	OpenDate string `flatfile:"4,10" json:"openDate" xml:"openDate"`
	Age      uint   `flatfile:"14,3" json:"age" xml:"age"`
}

type CustomerAddress struct {
	Address     string `flatfile:"1,15" json:"address" xml:"addressLine"`
	CountryCode string `flatfile:"16,2" json:"countryCode" xml:"countryCode"`
}
type CustomerRecord struct {
	//ffparser is one indexed, position starts at 1
	Demographics CustomerDemographic `flatfile:"1,16" json:"customerDemo" xml:"customerDemo"`
	Address      CustomerAddress     `flatfile:"17,17" json:"customerAddr" xml:"customerAddr"`
	PhoneNumbers [2]string           `flatfile:"34,10" json:"phoneNumbers" xml:"phoneNumbers>phoneNumber"`
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
			convertToJSON(fileRecord)
			convertToXML(fileRecord)
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

func convertToJSON(customerRecord *CustomerRecord) {
	custJSON, _ := json.MarshalIndent(customerRecord, "", "\t")
	fmt.Println(string(custJSON))
}

func convertToXML(customerRecord *CustomerRecord) {
	custXML, _ := xml.MarshalIndent(customerRecord, "", "\t")
	fmt.Println(string(custXML))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
