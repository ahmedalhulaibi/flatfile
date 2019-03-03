package main

import (
	"fmt"

	"github.com/ahmedalhulaibi/ffparser"
)

// To run this example:
// go run main.go

type CustomerDemographic struct {
	Name     string `ffp:"1,3"`
	OpenDate string `ffp:"4,10"`
	Age      uint   `ffp:"14,3"`
}

type CustomerAddress struct {
	Address     string `ffp:"1,15"`
	CountryCode string `ffp:"16,2"`
}

type CustomerRecord struct {
	//ffparser is one indexed, position starts at 1
	Demographics CustomerDemographic `ffp:"1,16"`
	Address      CustomerAddress     `ffp:"17,17"` //Nested struct. Define the address field range. Within the CustomerAddress struct, ffp tags should start at 1
}

func main() {
	data := []byte("AMY1900-01-01019123 FAKE STREETCA")

	fileRecord := &CustomerRecord{}
	ffparser.Examine(fileRecord)

	//unmarhsal text data to struct
	err := ffparser.Unmarshal(data, fileRecord, 0, 0)
	fmt.Printf("%v\n", fileRecord)

	if err != nil {
		fmt.Println(err)
	}
}
