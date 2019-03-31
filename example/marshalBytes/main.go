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
}

func main() {
	data := []byte("AMY1900-01-01019123 FAKE STREETCA")

	fileRecord := &CustomerRecord{}
	flatfile.Examine(fileRecord)

	//unmarhsal text data to struct
	err := flatfile.Unmarshal(data, fileRecord, 0, 0, false)
	fmt.Printf("%v\n", fileRecord)

	if err != nil {
		fmt.Println(err)
	}
}
