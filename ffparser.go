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

	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		//Get underlying type
		vType := reflect.TypeOf(v).Elem()

		//Only process if kind is Struct
		if vType.Kind() == reflect.Struct {
			//Dereference pointer to struct
			vStruct := reflect.ValueOf(v).Elem()

			//Loop through struct fields/properties
			for i := 0; i < vStruct.NumField(); i++ {

				//Get underlying type of field
				fieldType := vStruct.Field(i).Type()

				fieldTag, tagFlag := vType.Field(i).Tag.Lookup("ffp")
				if tagFlag {
					//split tag by comma to get position and length data
					params := strings.Split(fieldTag, ",")
					pos, poserr := strconv.Atoi(params[0])
					if poserr != nil {
						fmt.Println(poserr)
					}
					len, lenerr := strconv.Atoi(params[1])
					if lenerr != nil {
						fmt.Println(poserr)
					}

					//extract byte slice from byte data
					fieldData := data[pos-1 : pos-1+len]

					//fmt.Println(fieldType.String(), fieldType.Kind().String())

					assignBasedOnKind(fieldType.Kind(), vStruct.Field(i), fieldData)
				}
			}
		}

	}
	return nil
}

//assignBasedOnKind performs assignment of fieldData to field based on kind
func assignBasedOnKind(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	switch kind {
	case reflect.Uint8:
		var newFieldVal uint8
		buf := bytes.NewReader(fieldData)
		err := binary.Read(buf, binary.BigEndian, &newFieldVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(newFieldVal))
	case reflect.Uint32:
		newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 32)
		//fmt.Println(newFieldVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(uint32(newFieldVal)))
	case reflect.Uint64:
		//fmt.Println("uint64 found")
		newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 32)
		//fmt.Println(newFieldVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(newFieldVal))
	case reflect.String:
		field.Set(reflect.ValueOf(string(fieldData)))
	case reflect.Struct:
		Unmarshal(fieldData, field.Addr().Interface())
	case reflect.Ptr:
		//If pointer to struct
		if field.Elem().Kind() == reflect.Struct {
			//Unmarshal struct
			Unmarshal(fieldData, field.Interface())
		} else {
			fmt.Println(field.Elem())
			assignBasedOnKind(field.Elem().Kind(), field.Elem(), fieldData[:])
			fmt.Println(field.Elem())
		}
		//TODO: Pointer to other valid reflect.Kind
	case reflect.Slice, reflect.Array:
		//get underlying type, if struct then skip
		if field.Type().Elem().Kind() != reflect.Struct {
			// fmt.Println("SLICE or ARRAY found", field.Type().Elem().Kind())
			// fmt.Println(reflect.ValueOf(field))
			// fmt.Println(field.Index(0))
			for i := 0; i < field.Len(); i++ {
				//fmt.Println("sl element interface", field.Index(i))

				assignBasedOnKind(field.Type().Elem().Kind(), field.Index(i), fieldData[i:i+1])
			}
		}
	}
	return nil
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}
