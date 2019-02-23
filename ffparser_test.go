package ffparser

import (
	"reflect"
	"testing"
)

/*
uint8  : 0 to 255
uint16 : 0 to 65535
uint32 : 0 to 4294967295
uint64 : 0 to 18446744073709551615
int8   : -128 to 127
int16  : -32768 to 32767
int32  : -2147483648 to 2147483647
int64  : -9223372036854775808 to 9223372036854775807
*/

func TestBoolFalse(t *testing.T) {
	type BoolStruct struct {
		BoolFalse1 bool `ffp:"1,1"`
		BoolFalse2 bool `ffp:"2,1"`
		BoolFalse3 bool `ffp:"3,1"`
		BoolFalse4 bool `ffp:"4,5"`
		BoolFalse5 bool `ffp:"9,5"`
		BoolFalse6 bool `ffp:"14,5"`
	}

	testVal := &BoolStruct{}
	data := []byte("0fFfalseFalseFALSE")

	err := Unmarshal(data, testVal, 0)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	vStruct := reflect.ValueOf(testVal).Elem()

	for i := 0; i < vStruct.NumField(); i++ {
		t.Log(vStruct.Field(i).Bool())
		if vStruct.Field(i).Bool() {
			t.Fail()
		}
	}
}
func TestBoolTrue(t *testing.T) {
	type BoolStruct struct {
		BoolTrue1 bool `ffp:"1,1"`
		BoolTrue2 bool `ffp:"2,1"`
		BoolTrue3 bool `ffp:"3,1"`
		BoolTrue4 bool `ffp:"4,4"`
		BoolTrue5 bool `ffp:"8,4"`
		BoolTrue6 bool `ffp:"12,4"`
	}

	testVal := &BoolStruct{}
	data := []byte("1tTtrueTrueTRUE")

	err := Unmarshal(data, testVal, 0)

	if err != nil {
		t.Error(err)
	}
	vStruct := reflect.ValueOf(testVal).Elem()

	for i := 0; i < vStruct.NumField(); i++ {
		t.Log(vStruct.Field(i).Bool())
		if !vStruct.Field(i).Bool() {
			t.Fail()
		}
	}
}

func TestBoolErr(t *testing.T) {
	type BoolStruct struct {
		BoolTrue1 bool `ffp:"1,1"`
		BoolTrue2 bool `ffp:"2,1"`
		BoolTrue3 bool `ffp:"3,1"`
		BoolTrue4 bool `ffp:"4,4"`
		BoolTrue5 bool `ffp:"8,4"`
		BoolTrue6 bool `ffp:"12,4"`
	}

	testVal := &BoolStruct{}
	data := []byte("3aBerrrErrrERRR")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse bool")
	}
	t.Log(err)
}

func TestUint8(t *testing.T) {
	type Uint8Struct struct {
		Uint8One uint8 `ffp:"1,1"`
		Uint8Two uint8 `ffp:"2,3"`
	}

	testVal := &Uint8Struct{}
	data := []byte("1255")

	err := Unmarshal(data, testVal, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint8One != 1 || testVal.Uint8Two != 255 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestUint8InvalidSyntaxErr(t *testing.T) {
	type Uint8Struct struct {
		Uint8One uint8 `ffp:"1,1"`
	}

	testVal := &Uint8Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint8")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint8OutOfRangeErr(t *testing.T) {
	type Uint8Struct struct {
		Uint8One uint8 `ffp:"1,4"`
	}

	testVal := &Uint8Struct{}
	data := []byte("2555")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint8")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint16(t *testing.T) {
	type Uint16Struct struct {
		Uint16One uint16 `ffp:"1,1"`
		Uint16Two uint16 `ffp:"2,5"`
	}

	testVal := &Uint16Struct{}
	data := []byte("165535")

	err := Unmarshal(data, testVal, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint16One != 1 || testVal.Uint16Two != 65535 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestUint16InvalidSyntaxErr(t *testing.T) {
	type Uint16Struct struct {
		Uint16One uint16 `ffp:"1,1"`
	}

	testVal := &Uint16Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint16")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint16OutOfRangeErr(t *testing.T) {
	type Uint16Struct struct {
		Uint16One uint16 `ffp:"1,5"`
	}

	testVal := &Uint16Struct{}
	data := []byte("99999")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint16")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint32(t *testing.T) {
	type Uint32Struct struct {
		Uint32One uint32 `ffp:"1,1"`
		Uint32Two uint32 `ffp:"2,10"`
	}

	testVal := &Uint32Struct{}
	data := []byte("14294967295")

	err := Unmarshal(data, testVal, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint32One != 1 || testVal.Uint32Two != 4294967295 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestUint32InvalidSyntaxErr(t *testing.T) {
	type Uint32Struct struct {
		Uint32One uint32 `ffp:"1,1"`
	}

	testVal := &Uint32Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint32OutOfRangeErr(t *testing.T) {
	type Uint32Struct struct {
		Uint32One uint32 `ffp:"1,10"`
	}

	testVal := &Uint32Struct{}
	data := []byte("9999999999")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint64(t *testing.T) {
	type Uint64Struct struct {
		Uint64One uint64 `ffp:"1,1"`
		Uint64Two uint64 `ffp:"2,20"`
	}

	testVal := &Uint64Struct{}
	data := []byte("118446744073709551615")

	err := Unmarshal(data, testVal, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint64One != 1 || testVal.Uint64Two != 18446744073709551615 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestUint64InvalidSyntaxErr(t *testing.T) {
	type Uint64Struct struct {
		Uint64One uint64 `ffp:"1,1"`
	}

	testVal := &Uint64Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint64OutOfRangeErr(t *testing.T) {
	type Uint64Struct struct {
		Uint64One uint64 `ffp:"1,20"`
	}

	testVal := &Uint64Struct{}
	data := []byte("99999999999999999999")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt8(t *testing.T) {
	type Int8Struct struct {
		Int8One int8 `ffp:"1,4"`
		Int8Two int8 `ffp:"5,3"`
	}

	testVal := &Int8Struct{}
	data := []byte("-128127")

	err := Unmarshal(data, testVal, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int8One != -128 || testVal.Int8Two != 127 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt8InvalidSyntaxErr(t *testing.T) {
	type Int8Struct struct {
		Int8One int8 `ffp:"1,1"`
	}

	testVal := &Int8Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int8")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt8OutOfRangeErr(t *testing.T) {
	type Int8Struct struct {
		Int8One int8 `ffp:"1,4"`
	}

	testVal := &Int8Struct{}
	data := []byte("2555")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int8")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt16(t *testing.T) {
	type Int16Struct struct {
		Int16One int16 `ffp:"1,6"`
		Int16Two int16 `ffp:"7,5"`
	}

	testVal := &Int16Struct{}
	data := []byte("-3276832767")

	err := Unmarshal(data, testVal, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int16One != -32768 || testVal.Int16Two != 32767 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt16InvalidSyntaxErr(t *testing.T) {
	type Int16Struct struct {
		Int16One int16 `ffp:"1,1"`
	}

	testVal := &Int16Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int16")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt16OutOfRangeErr(t *testing.T) {
	type Int16Struct struct {
		Int16One int16 `ffp:"1,5"`
	}

	testVal := &Int16Struct{}
	data := []byte("99999")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int16")
	}
	t.Log(testVal)
	t.Log(err)
}
