package flatfile

import (
	"reflect"

	"github.com/pkg/errors"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*Unmarshal will read data and convert it into a struct based on a schema/map defined by struct tags

Struct tags are in the form `flatfile:"col,len"`. col and len should be integers > 0

startFieldIdx: index can be passed to indicate which struct field to start the unmarshal. Zero indexed.

numFieldsToUnmarshal: can be passed to indicate how many fields to unmarshal starting from startFieldIdx

isPartialUnmarshal: flag should be set to true if the data[0] is intended to be unmarshalled at startFieldIdx

If startFieldIdx == 0 and umFieldsToMarshal == 0 then Unmarshal will attempt to unmarshal all fields with an ffp tag

*/
func Unmarshal(data []byte, v interface{}, startFieldIdx int, numFieldsToUnmarshal int, isPartialUnmarshal bool) error {
	colOffset := 0
	//init ffpTag for later use
	ffpTag := &flatfileTag{}
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		//Get underlying type
		vType := reflect.TypeOf(v).Elem()

		//Only process if kind is Struct
		if vType.Kind() == reflect.Struct {
			//Dereference pointer to struct
			vStruct := reflect.ValueOf(v).Elem()
			maxField := 0
			if numFieldsToUnmarshal > 0 {
				maxField = min(startFieldIdx+numFieldsToUnmarshal, vStruct.NumField())
			} else {
				maxField = vStruct.NumField()
			}
			//Loop through struct fields/properties
			for i := startFieldIdx; i < maxField; i++ {

				//Get underlying type of field
				fieldType := vStruct.Field(i).Type()
				fieldTag, tagFlag := vType.Field(i).Tag.Lookup("flatfile")
				if tagFlag {

					tagParseErr := parseFlatfileTag(fieldTag, ffpTag)
					if tagParseErr != nil {
						return errors.Wrapf(tagParseErr, "flatfile.Unmarshal: Failed to parse field tag %s", fieldTag)
					}
					if ShouldUnmarshal(ffpTag, data) {
						//determine pos offset based on start index in case start index not 0 (1)
						if i == startFieldIdx && startFieldIdx > 0 && isPartialUnmarshal {
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
									return errors.Wrap(err, "flatfile.Unmarshal: Failed to unmarshal")
								}
							}
						}
					}
				}
			}
		}
		return nil
	}
	return errors.Errorf("flatfile.Unmarshal: Unmarshal not complete. %s is not a pointer", reflect.TypeOf(v))
}

//CalcNumFieldsToUnmarshal determines how many fields can be unmarshalled successfully
//This currently will not return an accurate result for overlapping fields
//For example:
//type Profile struct {
//		FirstName string `flatfile:"1,10"`
//		LastName  string `flatfile:"11,10"`
//		FullName  string `flatfile:"1,20"`
//}
//Expected output is that 3 fields can be unmarshalled successfully
//The result will be 2
//Another example:
//type Profile struct
//type Profile struct {
//		FirstName string `flatfile:"1,10"`
//		LastName  string `flatfile:"11,10"`
//		FullName  string `flatfile:"1,20"`
//		Random    string `flatfile:"7,9"`
//}
//This function would have to be redesigned to handle multiple scenarios of overlapping fields
func CalcNumFieldsToUnmarshal(data []byte, v interface{}, fieldOffset int) (int, []byte, error) {
	ffpTag := &flatfileTag{}
	dataLen := len(data)
	numFieldsToUnmarshal := 0
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

				fieldTag, tagFlag := vType.Field(i).Tag.Lookup("flatfile")
				if tagFlag {

					tagParseErr := parseFlatfileTag(fieldTag, ffpTag)
					if tagParseErr != nil {
						return 0, []byte(""), errors.Wrapf(tagParseErr, "flatfile.CalcNumFieldsToUnmarshal: Failed to parse field tag %s", fieldTag)
					}

					if ffpTag.occurs > 0 {
						cumulativeRecLength += ffpTag.length * ffpTag.occurs
					} else if fieldType.Kind() == reflect.Array {
						cumulativeRecLength += ffpTag.length * vStruct.Field(i).Len()
					} else {
						cumulativeRecLength += ffpTag.length
					}

					if cumulativeRecLength <= dataLen {
						numFieldsToUnmarshal++
					} else {
						remainder = data[(cumulativeRecLength - ffpTag.length):]
						break
					}
				}
			}
		}
		return numFieldsToUnmarshal, remainder, nil
	}
	return 0, []byte(""), errors.Errorf("flatfile.CalcNumFieldsToUnmarshal: CalcNumFieldsToUnmarshal not complete. %s is not a pointer", reflect.TypeOf(v))
}

//ShouldUnmarshal returns true if the condition
func ShouldUnmarshal(ffpTag *flatfileTag, data []byte) bool {
	if ffpTag.condChk {
		lowerBound := ffpTag.condCol - 1
		upperBound := lowerBound + ffpTag.condLen
		return string(data[lowerBound:upperBound]) == ffpTag.condVal
	}

	return true
}
