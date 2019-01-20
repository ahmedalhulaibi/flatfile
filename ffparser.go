package ffparser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)
/*Unmarshal maps a slice of bytes to a struct based on defined ffp struct tags 
Struct tags are in the form `ffp:"pos,len"` 
Example: 
type FileHeader struct {
	LogicalRecordTypeID byte   `ffp:"1,1"`
	LogicalRecordCount  uint32 `ffp:"2,9"`
	OriginatorID        string `ffp:"11,10"`
}
*/ 
func Unmarshal(data []byte, v interface{}) error {
	vType := reflect.TypeOf(v).Elem()

	vStruct := reflect.ValueOf(v).Elem()
	for i := 0; i < vStruct.NumField(); i++ {
		fieldType := vStruct.Field(i).Type()
		fieldTag, tagFlag := vType.Field(i).Tag.Lookup("ffp")
		if tagFlag {
			params := strings.Split(fieldTag, ",")
			pos, poserr := strconv.Atoi(params[0])
			if poserr != nil {
				fmt.Println(poserr)
			}
			len, lenerr := strconv.Atoi(params[1])
			if lenerr != nil {
				fmt.Println(poserr)
			}
			fieldData := data[pos-1 : pos-1+len]

			switch fieldType.String() {
			case "uint8":
				var newFieldVal uint8
				buf := bytes.NewReader(fieldData)
				err := binary.Read(buf, binary.BigEndian, &newFieldVal)
				if err != nil {
					return err
				}
				vStruct.Field(i).Set(reflect.ValueOf(newFieldVal))
			case "uint32":
				newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 32)
				//fmt.Println(newFieldVal)
				if err != nil {
					return err
				}
				vStruct.Field(i).Set(reflect.ValueOf(uint32(newFieldVal)))
			case "uint64":
				//fmt.Println("uint64 found")
				newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 32)
				//fmt.Println(newFieldVal)
				if err != nil {
					return err
				}
				vStruct.Field(i).Set(reflect.ValueOf(newFieldVal))
			case "[]uint8":
				//fmt.Println("[]uint8 found")
				vStruct.Field(i).Set(reflect.ValueOf(fieldData))
			case "string":
				vStruct.Field(i).Set(reflect.ValueOf(string(fieldData)))

			}
		}
	}

	return nil
}
