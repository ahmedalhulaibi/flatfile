package flatfile

import (
	"reflect"
	"strconv"
	"unicode/utf8"
	"unsafe"

	"github.com/pkg/errors"
)

//assignBasedOnKind performs assignment of fieldData to field based on kind
func assignBasedOnKind(kind reflect.Kind, field reflect.Value, fieldData []byte, ffpTag *flatfileTag) error {
	var err error
	err = nil
	switch kind {
	case reflect.Bool:
		err = assignBool(kind, field, fieldData)
	case reflect.Uint:
		err = assignUint(kind, field, fieldData)
	case reflect.Uint8:
		//check ffpTag.override == byte, meaning user wants to store the byte value itself
		if ffpTag.override == "byte" {
			err = assignByte(field, fieldData[0])
		} else {
			err = assignUint8(kind, field, fieldData)
		}

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
		//check ffpTag.override == rune, meaning user wants to store the rune value itself
		if ffpTag.override == "rune" {
			err = assignRune(field, fieldData)
		} else {
			err = assignInt32(kind, field, fieldData)
		}
	case reflect.Int64:
		err = assignInt64(kind, field, fieldData)
	case reflect.Float32:
		err = assignFloat32(kind, field, fieldData)
	case reflect.Float64:
		err = assignFloat64(kind, field, fieldData)
	case reflect.String:
		field.Set(reflect.ValueOf(string(fieldData)))
	case reflect.Struct:
		err = Unmarshal(fieldData, field.Addr().Interface(), 0, 0, false)
	case reflect.Ptr:
		//If pointer to struct
		if field.Elem().Kind() == reflect.Struct {
			//Unmarshal struct
			err = Unmarshal(fieldData, field.Interface(), 0, 0, false)
		} else {
			err = assignBasedOnKind(field.Elem().Kind(), field.Elem(), fieldData, ffpTag)
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
			err = errors.Errorf("flatfile.assignBasedOnKind: Occurs clause must be provided when using slice. `flatfile:\"col,len,occurs\"`")
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
	return errors.Wrap(err, "flatfile.assignBasedOnKind: AssignmentError")
}

func assignBool(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseBool(string(fieldData))
	//fmt.Println(newFieldVal)
	if err == nil {
		field.Set(reflect.ValueOf(newFieldVal))
	}

	return errors.Wrap(err, "flatfile.assignBool error")
}

func assignUint(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	var dummy uint
	//Determine bitness using Sizeof
	//this will return 1 for 8-bit, 2 for 16-bit, 4 for 32-bit, 8 for 64-bit. Multiply the result to get bitsize and convert to int for Parsing
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, int(unsafe.Sizeof(dummy)*8))
	if err == nil {
		field.Set(reflect.ValueOf(uint(newFieldVal)))
	}
	return errors.Wrapf(err, "flatfile.assignUint: Failed to assignUint %v ", field)
}

func assignUint8(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 8)
	if err == nil {
		field.Set(reflect.ValueOf(uint8(newFieldVal)))
	}
	return errors.Wrap(err, "flatfile.assignUint8 error")
}

func assignUint16(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 16)
	if err == nil {
		field.Set(reflect.ValueOf(uint16(newFieldVal)))
	}
	return errors.Wrap(err, "flatfile.assignUint16 error")
}

func assignUint32(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 32)
	if err == nil {
		field.Set(reflect.ValueOf(uint32(newFieldVal)))
	}
	return errors.Wrap(err, "flatfile.assignUint32 error")
}

func assignUint64(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseUint(string(fieldData), 10, 64)
	if err == nil {
		field.Set(reflect.ValueOf(newFieldVal))
	}
	return errors.Wrap(err, "flatfile.assignUint64 error")
}

func assignInt(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	var dummy int
	//Determine bitness using Sizeof
	//this will return 1 for 8-bit, 2 for 16-bit, 4 for 32-bit, 8 for 64-bit. Multiply the result to get bitsize and convert to int for Parsing
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, int(unsafe.Sizeof(dummy)*8))
	if err == nil {
		field.Set(reflect.ValueOf(int(newFieldVal)))
	}
	return errors.Wrapf(err, "flatfile.assignInt: Failed to assignInt %v ", field)
}

func assignInt8(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 8)
	if err == nil {
		field.Set(reflect.ValueOf(int8(newFieldVal)))
	}
	return errors.Wrap(err, "flatfile.assignInt8 error")
}

func assignInt16(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 16)
	if err == nil {
		field.Set(reflect.ValueOf(int16(newFieldVal)))
	}
	return errors.Wrap(err, "flatfile.assignInt16 error")
}

func assignInt32(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 32)
	if err == nil {
		field.Set(reflect.ValueOf(int32(newFieldVal)))
	}
	return errors.Wrap(err, "flatfile.assignInt32 error")
}

func assignInt64(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseInt(string(fieldData), 10, 64)
	if err == nil {
		field.Set(reflect.ValueOf(int64(newFieldVal)))
	}
	return errors.Wrap(err, "flatfile.assignInt64 error")
}

func assignFloat32(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseFloat(string(fieldData), 32)
	if err == nil {
		field.Set(reflect.ValueOf(float32(newFieldVal)))
	}
	return errors.Wrap(err, "flatfile.assignFloat32 error")
}

func assignFloat64(kind reflect.Kind, field reflect.Value, fieldData []byte) error {
	newFieldVal, err := strconv.ParseFloat(string(fieldData), 64)
	if err == nil {
		field.Set(reflect.ValueOf(newFieldVal))
	}
	return errors.Wrap(err, "flatfile.assignFloat64 error")
}

func assignByte(field reflect.Value, fieldData byte) error {
	field.Set(reflect.ValueOf(fieldData))
	return nil
}

func assignRune(field reflect.Value, fieldData []byte) error {
	newFieldVal, _ := utf8.DecodeRune(fieldData)
	if newFieldVal == utf8.RuneError {
		return errors.New("flatfile.assignRune error")
	}
	field.Set(reflect.ValueOf(newFieldVal))
	return nil
}
