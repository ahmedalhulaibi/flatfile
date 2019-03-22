package main

import (
	"fmt"

	"github.com/ahmedalhulaibi/ffparser"
)

type CustomerRecord struct {
	//ffparser is one indexed, position starts at 1
	Name        string `flatfile:"1,3"`
	OpenDate    string `flatfile:"4,10"`
	Age         uint   `flatfile:"14,3"`
	Address     string `flatfile:"17,15"`
	CountryCode string `flatfile:"32,2"`
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
