package ffparser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

type ffpTagType struct {
	pos    int
	length int
	occurs int
}

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
func Unmarshal(data []byte, v interface{}, posOffset uint) error {
	//init ffpTag for later use
	ffpTag := &ffpTagType{}
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
					// params := strings.Split(fieldTag, ",")
					// pos, poserr := strconv.Atoi(params[0])
					// if poserr != nil {
					// 	fmt.Println(poserr)
					// }
					// len, lenerr := strconv.Atoi(params[1])
					// if lenerr != nil {
					// 	fmt.Println(poserr)
					// }
					tagParseErr := parseFfpTag(fieldTag, ffpTag)
					if tagParseErr != nil {
						return fmt.Errorf("ffparser: Failed to parse field tag %s:\n\t%s", fieldTag, tagParseErr)
					}
					//extract byte slice from byte data
					pos := ffpTag.pos - 1
					fieldData := data[pos : pos+ffpTag.length]

					//fmt.Println(fieldType.String(), fieldType.Kind().String())

					err := assignBasedOnKind(fieldType.Kind(), vStruct.Field(i), fieldData, ffpTag)
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

//parseFfpTag parses an ffp struct tag on a field
//Tags are expected to be in the form:
// pos,len,occurs
// where pos is an int > 0
//		 len is an int
func parseFfpTag(fieldTag string, ffpTag *ffpTagType) error {

	//split tag by comma to get position and length data
	params := strings.Split(fieldTag, ",")
	//position and length parameters must be provided
	//
	if len(params) < 2 {
		return fmt.Errorf("ffparser: Not enough ffp tag params provided.\nPosition and length parameters must be provided.\nMust be in form `ffp:\"pos,len\"`")
	}

	pos, poserr := strconv.Atoi(params[0])
	if poserr != nil {
		return fmt.Errorf("ffparser: Error parsing position parameter\n%s", poserr)
	}

	if pos < 1 {
		return fmt.Errorf("ffparser: Out of range error. Position parameter cannot be less than 1. Please note position is 1-indexed not zero")
	}

	ffpTag.pos = pos

	length, lenerr := strconv.Atoi(params[1])
	if lenerr != nil {
		return fmt.Errorf("ffparser: Error parsing length parameter\n%s", lenerr)
	}

	if length < 1 {
		return fmt.Errorf("ffparser: Out of range error. Length parameter cannot be less than 1")
	}

	ffpTag.length = length

	if len(params) > 2 {
		occurs, occerr := strconv.Atoi(params[2])
		if occerr != nil {
			return fmt.Errorf("ffparser: Error parsing occurs parameter\n%s", occerr)
		}

		if occurs < 2 {
			return fmt.Errorf("ffparser: Out of range error. Occurs parameter cannot be less than 2")
		}

		ffpTag.occurs = occurs
	}

	return nil
}

//assignBasedOnKind performs assignment of fieldData to field based on kind
func assignBasedOnKind(kind reflect.Kind, field reflect.Value, fieldData []byte, ffpTag *ffpTagType) error {
	var err error
	err = nil
	switch kind {
	case reflect.Bool:
		err = assignBool(kind, field, fieldData)
	case reflect.Uint:
		err = assignUint(kind, field, fieldData)
	case reflect.Uint8:
		err = assignUint8(kind, field, fieldData)
	case reflect.Uint16:
		err = assignUint16(kind, field, fieldData)
	case reflect.Uint32:
		err = assignUint32(kind, field, fieldData)
	case reflect.Uint64:
		err = assignUint64(kind, field, fieldData)
	case reflect.Int:
		err = assignInt(kind, field, fieldData)
	case reflect.Int8:
		err = assignInt8(kind, field, fieldData)
	case reflect.Int16:
		err = assignInt16(kind, field, fieldData)
	case reflect.Int32:
		err = assignInt32(kind, field, fieldData)
	case reflect.Int64:
		err = assignInt64(kind, field, fieldData)
	case reflect.Float32:
		err = assignFloat32(kind, field, fieldData)
	case reflect.Float64:
		err = assignFloat64(kind, field, fieldData)
	case reflect.String:
		field.Set(reflect.ValueOf(string(fieldData)))
	case reflect.Struct:
		err = Unmarshal(fieldData, field.Addr().Interface(), 0)
	case reflect.Ptr:
		//If pointer to struct
		if field.Elem().Kind() == reflect.Struct {
			//Unmarshal struct
			err = Unmarshal(fieldData, field.Interface(), 0)
		} else {
			fmt.Println(field.Elem())
			err = assignBasedOnKind(field.Elem().Kind(), field.Elem(), fieldData[:], ffpTag)
			fmt.Println(field.Elem())
		}
	case reflect.Array:
		for i := 0; i < field.Len(); i++ {
			//fmt.Println("sl element interface", field.Index(i))
			lowerBound := i * ffpTag.length
			upperBound := lowerBound + ffpTag.length
			assignBasedOnKind(field.Type().Elem().Kind(), field.Index(i), fieldData[lowerBound:upperBound], ffpTag)
		}
	case reflect.Slice:
		if ffpTag.occurs < 1 {
			err = fmt.Errorf("ffparser: Occurs clause must be provided when using slice. `ffp:\"pos,len,occurs\"`")
		}
		//make slice of length ffpTag.occurs to avoid index out of range err
		field.Set(reflect.MakeSlice(field.Type(), ffpTag.occurs, ffpTag.occurs))
		for i := 0; i < ffpTag.occurs; i++ {
			//fmt.Println("sl element interface", field.Index(i))
			lowerBound := i * ffpTag.length
			upperBound := lowerBound + ffpTag.length
			assignBasedOnKind(field.Type().Elem().Kind(), field.Index(i), fieldData[lowerBound:upperBound], ffpTag)
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

func assignUint(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	//Determine bitness using Sizeof
	var dummy uint
	switch unsafe.Sizeof(dummy) {
	case 1:
		newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 8)
		if err == nil {
			field.Set(reflect.ValueOf(uint(newFieldVal)))
		}
		return err
	case 2:
		newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 16)
		if err == nil {
			field.Set(reflect.ValueOf(uint(newFieldVal)))
		}
		return err
	case 4:
		newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 32)
		if err == nil {
			field.Set(reflect.ValueOf(uint(newFieldVal)))
		}
		return err
	case 8:
		newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 64)
		if err == nil {
			field.Set(reflect.ValueOf(uint(newFieldVal)))
		}
		return err
	}
	return fmt.Errorf("ffparser: Failed to assignUint %v ", field)
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

func assignInt(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	//Determine bitness using Sizeof
	var dummy int
	switch unsafe.Sizeof(dummy) {
	case 1:
		newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 8)
		if err == nil {
			field.Set(reflect.ValueOf(int(newFieldVal)))
		}
		return err
	case 2:
		newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 16)
		if err == nil {
			field.Set(reflect.ValueOf(int(newFieldVal)))
		}
		return err
	case 4:
		newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 32)
		if err == nil {
			field.Set(reflect.ValueOf(int(newFieldVal)))
		}
		return err
	case 8:
		newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 64)
		if err == nil {
			field.Set(reflect.ValueOf(int(newFieldVal)))
		}
		return err
	}
	return fmt.Errorf("ffparser: Failed to assignInt %v ", field)
}

func assignInt8(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 8)
	if err == nil {
		field.Set(reflect.ValueOf(int8(newFieldVal)))
	}
	return err
}

func assignInt16(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 16)
	if err == nil {
		field.Set(reflect.ValueOf(int16(newFieldVal)))
	}
	return err
}

func assignInt32(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 32)
	if err == nil {
		field.Set(reflect.ValueOf(int32(newFieldVal)))
	}
	return err
}

func assignInt64(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 64)
	if err == nil {
		field.Set(reflect.ValueOf(int64(newFieldVal)))
	}
	return err
}

func assignFloat32(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseFloat(string(fieldData), 32)
	if err == nil {
		field.Set(reflect.ValueOf(float32(newFieldVal)))
	}
	return err
}

func assignFloat64(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseFloat(string(fieldData), 64)
	if err == nil {
		field.Set(reflect.ValueOf(newFieldVal))
	}
	return err
}

// Examine traverses all elements of a type and uses the reflect pkg to print type and kind
func Examine(v interface{}) {
	examiner(reflect.TypeOf(v), 0)
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
