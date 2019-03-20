package main

import (
	"fmt"

	"github.com/ahmedalhulaibi/ffparser"
)

type CustomerRecord struct {
	//ffparser is one indexed, position starts at 1
	//override option has to be supplied to convert byte/rune correctly
	FirstInitial byte   `ffp:"col=1,len=1,override=byte"`
	Name         string `ffp:"col=1,len=3"`
	OpenDate     string `ffp:"column=4,length=10"`
	Age          uint   `ffp:"col=14,3"`
	//you can mix named with unnamed, but keep in mind unnamed options must be in the correct position (col,len,occurs,override)
	Address string `ffp:"17,len=15"`
	//named options can be in any order
	CountryCode string `ffp:"len=2,col=32"`
	//The below tag is in the form "col,len,occurences"
	//The phone numbers start at position 34 (one indexed)
	//The phone numbers are each number 10 bytes long
	//There are 2 occurences of phone numbers total
	//For clarity the second phone number will be read in from pos 44 (one indexed)
	PhoneNumbers []string `ffp:"col=34,len=10,occurs=2"`
}

func main() {
	data := []byte("AMY1900-01-01019123 FAKE STREETCA41611122229053334444")

	fileHeader := &CustomerRecord{}
	ffparser.Examine(fileHeader)

	err := ffparser.Unmarshal(data, fileHeader, 0, 0)
	fmt.Printf("%v\n", fileHeader)

	if err != nil {
		fmt.Println(err)
	}
}
