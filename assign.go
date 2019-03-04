package ffparser

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

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
		err = Unmarshal(fieldData, field.Addr().Interface(), 0, 0)
	case reflect.Ptr:
		//If pointer to struct
		if field.Elem().Kind() == reflect.Struct {
			//Unmarshal struct
			err = Unmarshal(fieldData, field.Interface(), 0, 0)
		} else {
			err = assignBasedOnKind(field.Elem().Kind(), field.Elem(), fieldData[:], ffpTag)
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
