package main

import (
	"fmt"

	"github.com/ahmedalhulaibi/flatfile"
)

// To run this example:
// go run main.go

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
	Demographics CustomerDemographic `flatfile:"1,16"`
	Address      CustomerAddress     `flatfile:"17,17"` //Nested struct. Define the address field range. Within the CustomerAddress struct, ffp tags should start at 1
}

func main() {
	data := []byte("AMY1900-01-01019123 FAKE STREETCA")

	fileRecord := &CustomerRecord{}
	flatfile.Examine(fileRecord)

	//unmarhsal text data to struct
	err := flatfile.Unmarshal(data, fileRecord, 0, 0)
	fmt.Printf("%v\n", fileRecord)

	if err != nil {
		fmt.Println(err)
	}
}
