package ffparser

import (
	"fmt"
	"reflect"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*Unmarshal will read data and convert it into a struct based on a schema/map defined by struct tags

Struct tags are in the form `ffp:"col,len"`. col and len should be integers > 0

startFieldIdx: index can be passed to indicate which struct field to start the unmarshal

numFieldsToMarshal: can be passed to indicate how many fields to unmarshal starting from startFieldIdx


If startFieldIdx == 0 and umFieldsToMarshal == 0 then Unmarshal will attempt to marshal all fields with an ffp tag

*/
func Unmarshal(data []byte, v interface{}, startFieldIdx int, numFieldsToMarshal int) error {
	colOffset := 0
	//init ffpTag for later use
	ffpTag := &ffpTagType{}
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		//Get underlying type
		vType := reflect.TypeOf(v).Elem()

		//Only process if kind is Struct
		if vType.Kind() == reflect.Struct {
			//Dereference pointer to struct
			vStruct := reflect.ValueOf(v).Elem()
			maxField := 0
			if numFieldsToMarshal > 0 {
				maxField = min(startFieldIdx+numFieldsToMarshal, vStruct.NumField())
			} else {
				maxField = vStruct.NumField()
			}
			//Loop through struct fields/properties
			for i := startFieldIdx; i < maxField; i++ {

				//Get underlying type of field
				fieldType := vStruct.Field(i).Type()
				fieldTag, tagFlag := vType.Field(i).Tag.Lookup("ffp")
				if tagFlag {

					tagParseErr := parseFfpTag(fieldTag, ffpTag)
					if tagParseErr != nil {
						return fmt.Errorf("ffparser: Failed to parse field tag %s:\n\t%s", fieldTag, tagParseErr)
					}
					//determine pos offset based on start index in case start index not 1
					if i == startFieldIdx {
						colOffset = ffpTag.col - 1
					}

					//determine if the current field is in range of the posOffset passed
					if ffpTag.col > colOffset {
						//extract byte slice from byte data
						lowerBound := ffpTag.col - 1 - colOffset
						upperBound := lowerBound + ffpTag.length
						//and check that pos does not exceed length of bytes to prevent attempting to parse nulls
						if lowerBound < len(data) {
							fieldData := data[lowerBound:upperBound]

							err := assignBasedOnKind(fieldType.Kind(), vStruct.Field(i), fieldData, ffpTag)
							if err != nil {
								return fmt.Errorf("ffparser: Failed to marshal.\n%s", err)
							}
						}
					}
				}
			}
		}
		return nil
	}
	return fmt.Errorf("ffparser: Unmarshal not complete. %s is not a pointer", reflect.TypeOf(v))
}

//CalcNumFieldsToMarshal determines how many fields can be marshalled successfully
//This currently will not return an accurate result for overlapping fields
//For example:
//type Profile struct {
//		FirstName string `ffp:"1,10"`
//		LastName  string `ffp:"11,10"`
//		FullName  string `ffp:"1,20"`
//}
//Expected output is that 3 fields can be marshalled successfully
//The result will be 2
//Another example:
//type Profile struct
//type Profile struct {
//		FirstName string `ffp:"1,10"`
//		LastName  string `ffp:"11,10"`
//		FullName  string `ffp:"1,20"`
//		Random    string `ffp:"7,9"`
//}
//This function would have to be redesigned to handle multiple scenarios of overlapping fields
func CalcNumFieldsToMarshal(data []byte, v interface{}, fieldOffset int) (int, []byte, error) {
	ffpTag := &ffpTagType{}
	dataLen := len(data)
	numFieldsToMarshal := 0
	var remainder []byte
	cumulativeRecLength := 0
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		//Get underlying type
		vType := reflect.TypeOf(v).Elem()

		//Only process if kind is Struct
		if vType.Kind() == reflect.Struct {
			//Dereference pointer to struct
			vStruct := reflect.ValueOf(v).Elem()

			//Loop through struct fields/properties
			for i := fieldOffset; i < vStruct.NumField(); i++ {

				//Get underlying type of field
				fieldType := vStruct.Field(i).Type()

				fieldTag, tagFlag := vType.Field(i).Tag.Lookup("ffp")
				if tagFlag {

					tagParseErr := parseFfpTag(fieldTag, ffpTag)
					if tagParseErr != nil {
						return 0, []byte(""), fmt.Errorf("ffparser: Failed to parse field tag %s:\n\t%s", fieldTag, tagParseErr)
					}

					if ffpTag.occurs > 0 {
						cumulativeRecLength += ffpTag.length * ffpTag.occurs
					} else if fieldType.Kind() == reflect.Array {
						cumulativeRecLength += ffpTag.length * vStruct.Field(i).Len()
					} else {
						cumulativeRecLength += ffpTag.length
					}

					if cumulativeRecLength <= dataLen {
						numFieldsToMarshal++
					} else {
						remainder = data[(cumulativeRecLength - ffpTag.length):]
						break
					}
				}
			}
		}
		return numFieldsToMarshal, remainder, nil
	}
	return 0, []byte(""), fmt.Errorf("ffparser: Unmarshal not complete. %s is not a pointer", reflect.TypeOf(v))
}
