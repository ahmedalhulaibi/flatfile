package ffparser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*Unmarshal will read data and convert it into a struct based on a schema/map defined by struct tags
Struct tags are in the form `ffp:"pos,len"`. An offset can be pass to the function when reading long lines of data.
The offset will be added to pos.
Example:
type FileHeader struct {
	LogicalRecordTypeID byte   `ffp:"1,1"`
	LogicalRecordCount  uint32 `ffp:"2,9"`
	OriginatorID        string `ffp:"11,10"`
}
data: contains the data that will be mapped to struct properties
v: must be a pointer to a struct
posOffset:
*/
func Unmarshal(data []byte, v interface{}, posOffset int) error {

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

					err := assignBasedOnKind(fieldType.Kind(), vStruct.Field(i), fieldData)
					if err != nil {
						return fmt.Errorf("ffparser: Failed to marshal.\n%s", err)
					}
				}
			}
		}
		return nil
	}
	return fmt.Errorf("ffparser: Unmarshal not complete. %s is not a pointer", reflect.TypeOf(v))
}

//assignBasedOnKind performs assignment of fieldData to field based on kind
func assignBasedOnKind(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	var err error
	err = nil
	switch kind {
	case reflect.Bool:
		err = assignBool(kind, field, fieldData)
	case reflect.Uint8:
		err = assignUint8(kind, field, fieldData)
	case reflect.Uint16:
		err = assignUint16(kind, field, fieldData)
	case reflect.Uint32:
		err = assignUint32(kind, field, fieldData)
	case reflect.Uint64:
		err = assignUint64(kind, field, fieldData)
	case reflect.String:
		field.Set(reflect.ValueOf(string(fieldData)))
	case reflect.Struct:
		Unmarshal(fieldData, field.Addr().Interface(), 0)
	case reflect.Ptr:
		//If pointer to struct
		if field.Elem().Kind() == reflect.Struct {
			//Unmarshal struct
			Unmarshal(fieldData, field.Interface(), 0)
		} else {
			fmt.Println(field.Elem())
			assignBasedOnKind(field.Elem().Kind(), field.Elem(), fieldData[:])
			fmt.Println(field.Elem())
		}
	case reflect.Slice, reflect.Array:
		//get underlying type, if not struct (slice, array) then skip
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
	return err
}

func assignBool(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseBool(string(fieldData))
	//fmt.Println(newFieldVal)
	if err == nil {
		field.Set(reflect.ValueOf(newFieldVal))
	}
	return err
}

func assignUint8(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 8)
	if err == nil {
		field.Set(reflect.ValueOf(uint8(newFieldVal)))
	}
	return err
}

func assignUint16(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 16)
	if err == nil {
		field.Set(reflect.ValueOf(uint16(newFieldVal)))
	}
	return err
}

func assignUint32(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 32)
	if err == nil {
		field.Set(reflect.ValueOf(uint32(newFieldVal)))
	}
	return err
}

func assignUint64(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 64)
	if err == nil {
		field.Set(reflect.ValueOf(newFieldVal))
	}
	return err
}

// Below code is sourced from Jon Bodner's blog: https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
// Direct link to Gist: https://gist.github.com/jonbodner/1727d0825d73541db8d6fcb859515735
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
