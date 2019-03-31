package main

import (
	"fmt"

	"github.com/ahmedalhulaibi/flatfile"
)

type CustomerRecord struct {
	//flatfile is one indexed, position starts at 1
	//override option has to be supplied to convert byte/rune correctly
	FirstInitial     byte   `flatfile:"col=1,len=1,override=byte"`
	FirstInitialRune rune   `flatfile:"col=1,len=1,override=rune"`
	Name             string `flatfile:"col=1,len=3"`
	OpenDate         string `flatfile:"column=4,length=10"`
	Age              uint   `flatfile:"col=14,3"`
	//you can mix named with unnamed, but keep in mind unnamed options must be in the correct position (col,len,occurs,override)
	Address string `flatfile:"17,len=15"`
	//named options can be in any order
	CountryCode string `flatfile:"len=2,col=32"`
	//The below tag is in the form "col,len,occurences"
	//The phone numbers start at position 34 (one indexed)
	//The phone numbers are each number 10 bytes long
	//There are 2 occurences of phone numbers total
	//For clarity the second phone number will be read in from pos 44 (one indexed)
	PhoneNumbers []string `flatfile:"col=34,len=10,occurs=2"`
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
