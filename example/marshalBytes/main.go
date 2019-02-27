package main

import (
	"fmt"

	"github.com/ahmedalhulaibi/ffparser"
)

type CustomerRecord struct {
	//ffparser is one indexed, position starts at 1
	Name        string `ffp:"1,3"`
	OpenDate    string `ffp:"4,10"`
	Age         uint   `ffp:"14,3"`
	Address     string `ffp:"17,15"`
	CountryCode string `ffp:"32,2"`
}

func main() {
	data := []byte("AMY1900-01-01019123 FAKE STREETCA")

	fileHeader := &CustomerRecord{}
	ffparser.Examine(fileHeader)

	err := ffparser.Unmarshal(data, fileHeader, 0)
	fmt.Printf("%v\n", fileHeader)

	if err != nil {
		fmt.Println(err)
	}
}
