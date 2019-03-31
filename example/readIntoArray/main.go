package main

import (
	"fmt"

	"github.com/ahmedalhulaibi/flatfile"
)

type CustomerRecord struct {
	//flatfile is one indexed, position starts at 1
	Name        string `flatfile:"1,3"`
	OpenDate    string `flatfile:"4,10"`
	Age         uint   `flatfile:"14,3"`
	Address     string `flatfile:"17,15"`
	CountryCode string `flatfile:"32,2"`
	//The below tag is in the form "col,len"
	//The phone numbers start at position 34 (one indexed)
	//The phone numbers are each number 10 bytes long
	//There are 2 phone numbers total
	//For clarity the second phone number will be read in from pos 44
	PhoneNumbers [2]string `flatfile:"34,10"`
}

func main() {
	data := []byte("AMY1900-01-01019123 FAKE STREETCA41611122229053334444")

	fileRecord := &CustomerRecord{}
	flatfile.Examine(fileRecord)

	err := flatfile.Unmarshal(data, fileRecord, 0, 0, false)
	fmt.Printf("%v\n", fileRecord)

	if err != nil {
		fmt.Println(err)
	}
}
