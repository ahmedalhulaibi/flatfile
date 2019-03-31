package flatfile

import (
	"bytes"
	"fmt"
	"math"
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

func TestBoolFalse_Unmarshal(t *testing.T) {
	type BoolStruct struct {
		BoolFalse1 bool `flatfile:"1,1"`
		BoolFalse2 bool `flatfile:"2,1"`
		BoolFalse3 bool `flatfile:"3,1"`
		BoolFalse4 bool `flatfile:"4,5"`
		BoolFalse5 bool `flatfile:"9,5"`
		BoolFalse6 bool `flatfile:"14,5"`
	}

	testVal := &BoolStruct{}
	data := []byte("0fFfalseFalseFALSE")

	err := Unmarshal(data, testVal, 0, 0, false)

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
func TestBoolTrue_Unmarshal(t *testing.T) {
	type BoolStruct struct {
		BoolTrue1 bool `flatfile:"1,1"`
		BoolTrue2 bool `flatfile:"2,1"`
		BoolTrue3 bool `flatfile:"3,1"`
		BoolTrue4 bool `flatfile:"4,4"`
		BoolTrue5 bool `flatfile:"8,4"`
		BoolTrue6 bool `flatfile:"12,4"`
	}

	testVal := &BoolStruct{}
	data := []byte("1tTtrueTrueTRUE")

	err := Unmarshal(data, testVal, 0, 0, false)

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

func TestBoolErr_Unmarshal(t *testing.T) {
	type BoolStruct struct {
		BoolTrue1 bool `flatfile:"1,1"`
		BoolTrue2 bool `flatfile:"2,1"`
		BoolTrue3 bool `flatfile:"3,1"`
		BoolTrue4 bool `flatfile:"4,4"`
		BoolTrue5 bool `flatfile:"8,4"`
		BoolTrue6 bool `flatfile:"12,4"`
	}

	testVal := &BoolStruct{}
	data := []byte("3aBerrrErrrERRR")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse bool")
	}
	t.Log(err)
}

func TestUint8_Unmarshal(t *testing.T) {
	type Uint8Struct struct {
		Uint8One uint8 `flatfile:"1,1"`
		Uint8Two uint8 `flatfile:"2,3"`
	}

	testVal := &Uint8Struct{}
	data := []byte("1255")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint8One != 1 || testVal.Uint8Two != 255 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestUint8InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Uint8Struct struct {
		Uint8One uint8 `flatfile:"1,1"`
	}

	testVal := &Uint8Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint8")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint8OutOfRangeErr_Unmarshal(t *testing.T) {
	type Uint8Struct struct {
		Uint8One uint8 `flatfile:"1,4"`
	}

	testVal := &Uint8Struct{}
	data := []byte("2555")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint8")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint16_Unmarshal(t *testing.T) {
	type Uint16Struct struct {
		Uint16One uint16 `flatfile:"1,1"`
		Uint16Two uint16 `flatfile:"2,5"`
	}

	testVal := &Uint16Struct{}
	data := []byte("165535")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint16One != 1 || testVal.Uint16Two != 65535 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestUint16InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Uint16Struct struct {
		Uint16One uint16 `flatfile:"1,1"`
	}

	testVal := &Uint16Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint16")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint16OutOfRangeErr_Unmarshal(t *testing.T) {
	type Uint16Struct struct {
		Uint16One uint16 `flatfile:"1,5"`
	}

	testVal := &Uint16Struct{}
	data := []byte("99999")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint16")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint32_Unmarshal(t *testing.T) {
	type Uint32Struct struct {
		Uint32One uint32 `flatfile:"1,1"`
		Uint32Two uint32 `flatfile:"2,10"`
	}

	testVal := &Uint32Struct{}
	data := []byte("14294967295")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint32One != 1 || testVal.Uint32Two != 4294967295 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestUint32InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Uint32Struct struct {
		Uint32One uint32 `flatfile:"1,1"`
	}

	testVal := &Uint32Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint32OutOfRangeErr_Unmarshal(t *testing.T) {
	type Uint32Struct struct {
		Uint32One uint32 `flatfile:"1,10"`
	}

	testVal := &Uint32Struct{}
	data := []byte("9999999999")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint64_Unmarshal(t *testing.T) {
	type Uint64Struct struct {
		Uint64One uint64 `flatfile:"1,1"`
		Uint64Two uint64 `flatfile:"2,20"`
	}

	testVal := &Uint64Struct{}
	data := []byte("118446744073709551615")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint64One != 1 || testVal.Uint64Two != 18446744073709551615 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestUint64InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Uint64Struct struct {
		Uint64One uint64 `flatfile:"1,1"`
	}

	testVal := &Uint64Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint64OutOfRangeErr_Unmarshal(t *testing.T) {
	type Uint64Struct struct {
		Uint64One uint64 `flatfile:"1,20"`
	}

	testVal := &Uint64Struct{}
	data := []byte("99999999999999999999")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse uint64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestUint_Unmarshal(t *testing.T) {
	type UintStruct struct {
		Uint8val  uint `flatfile:"1,1"`
		Uint16val uint `flatfile:"2,5"`
		Uint32val uint `flatfile:"2,10"`
		Uint64val uint `flatfile:"2,20"`
	}

	testVal := &UintStruct{}
	data := []byte("118446744073709551615")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Uint8val != 1 || testVal.Uint16val != 18446 || testVal.Uint32val != 1844674407 || testVal.Uint64val != 18446744073709551615 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt8_Unmarshal(t *testing.T) {
	type Int8Struct struct {
		Int8One int8 `flatfile:"1,4"`
		Int8Two int8 `flatfile:"5,3"`
	}

	testVal := &Int8Struct{}
	data := []byte("-128127")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int8One != -128 || testVal.Int8Two != 127 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt8InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Int8Struct struct {
		Int8One int8 `flatfile:"1,1"`
	}

	testVal := &Int8Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int8")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt8OutOfRangeErr_Unmarshal(t *testing.T) {
	type Int8Struct struct {
		Int8One int8 `flatfile:"1,4"`
	}

	testVal := &Int8Struct{}
	data := []byte("2555")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int8")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt16_Unmarshal(t *testing.T) {
	type Int16Struct struct {
		Int16One int16 `flatfile:"1,6"`
		Int16Two int16 `flatfile:"7,5"`
	}

	testVal := &Int16Struct{}
	data := []byte("-3276832767")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int16One != -32768 || testVal.Int16Two != 32767 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt16InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Int16Struct struct {
		Int16One int16 `flatfile:"1,1"`
	}

	testVal := &Int16Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int16")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt16OutOfRangeErr_Unmarshal(t *testing.T) {
	type Int16Struct struct {
		Int16One int16 `flatfile:"1,5"`
	}

	testVal := &Int16Struct{}
	data := []byte("99999")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int16")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt32_Unmarshal(t *testing.T) {
	type Int32Struct struct {
		Int32One int32 `flatfile:"1,11"`
		Int32Two int32 `flatfile:"12,10"`
	}

	testVal := &Int32Struct{}
	data := []byte("-21474836482147483647")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int32One != -2147483648 || testVal.Int32Two != 2147483647 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt32InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Int32Struct struct {
		Int32One int32 `flatfile:"1,1"`
	}

	testVal := &Int32Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt32OutOfRangeErr_Unmarshal(t *testing.T) {
	type Int32Struct struct {
		Int32One int32 `flatfile:"1,10"`
	}

	testVal := &Int32Struct{}
	data := []byte("9999999999")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt64_Unmarshal(t *testing.T) {
	type Int64Struct struct {
		Int64One int64 `flatfile:"1,20"`
		Int64Two int64 `flatfile:"21,19"`
	}

	testVal := &Int64Struct{}
	data := []byte("-92233720368547758089223372036854775807")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int64One != -9223372036854775808 || testVal.Int64Two != 9223372036854775807 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt64InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Int64Struct struct {
		Int64One int64 `flatfile:"1,1"`
	}

	testVal := &Int64Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt64OutOfRangeErr_Unmarshal(t *testing.T) {
	type Int64Struct struct {
		Int64One int64 `flatfile:"1,19"`
	}

	testVal := &Int64Struct{}
	data := []byte("9999999999999999999")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt_Unmarshal(t *testing.T) {
	type IntStruct struct {
		Int8val  int `flatfile:"1,1"`
		Int16val int `flatfile:"2,5"`
		Int32val int `flatfile:"2,9"`
		Int64val int `flatfile:"1,19"`
	}

	testVal := &IntStruct{}
	data := []byte("9223372036854775807")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int8val != 9 || testVal.Int16val != 22337 || testVal.Int32val != 223372036 || testVal.Int64val != 9223372036854775807 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestFloat32_Unmarshal(t *testing.T) {
	type Float32Struct struct {
		Float32One float32 `flatfile:"1,22"`
		Float32Two float32 `flatfile:"23,21"`
	}

	testVal := &Float32Struct{}
	data := []byte("3.4028234663852886e+381.401298464324817e-45")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Float32One != math.MaxFloat32 || testVal.Float32Two != math.SmallestNonzeroFloat32 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestFloat32InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Float32Struct struct {
		Float32One float32 `flatfile:"1,1"`
	}

	testVal := &Float32Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse Float32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFloat32OutOfRangeErr_Unmarshal(t *testing.T) {
	type Float32Struct struct {
		Float32One float32 `flatfile:"1,40"`
	}

	testVal := &Float32Struct{}
	data := []byte("9999999999999999999999999999999999999999")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse Float32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFloat64_Unmarshal(t *testing.T) {
	type Float64Struct struct {
		Float64One float64 `flatfile:"1,23"`
		Float64Two float64 `flatfile:"24,6"`
	}

	testVal := &Float64Struct{}
	data := []byte("1.7976931348623157e+3085e-324")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Error(err)
	}

	if testVal.Float64One != math.MaxFloat64 || testVal.Float64Two != math.SmallestNonzeroFloat64 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestFloat64InvalidSyntaxErr_Unmarshal(t *testing.T) {
	type Float64Struct struct {
		Float64One float64 `flatfile:"1,1"`
	}

	testVal := &Float64Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse Float64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFloat64OutOfRangeErr_Unmarshal(t *testing.T) {
	type Float64Struct struct {
		Float64One float64 `flatfile:"1,23"`
	}

	testVal := &Float64Struct{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse Float64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParsePosSyntaxErr_Unmarshal(t *testing.T) {
	type FfpTest struct {
		TestVal string `flatfile:"asdf,1"`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return syntax error when failing to parse position param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseLenSyntaxErr_Unmarshal(t *testing.T) {
	type FfpTest struct {
		TestVal string `flatfile:"1,asdf"`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return syntax error when failing to parse length param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParsePosRangeErr_Unmarshal(t *testing.T) {
	type FfpTest struct {
		TestVal string `flatfile:"-1,10"`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return out of range error when failing to parse position param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseLenRangeErr_Unmarshal(t *testing.T) {
	type FfpTest struct {
		TestVal string `flatfile:"1,-1"`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return out of range error when failing to parse length param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseMissingParamErr_Unmarshal(t *testing.T) {
	type FfpTest struct {
		TestVal string `flatfile:""`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return missing parameter error")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestArrayParse(t *testing.T) {
	type FfpTest struct {
		TestVal [4]int     `flatfile:"1,2"`
		Names   [10]string `flatfile:"9,3"`
	}

	testVal := &FfpTest{}
	expectedVal := [4]int{11, 22, 33, 44}

	expectedNames := [10]string{"AMY", "BOB", "CAM", "DAN", "EDD", "FAE", "GUY", "HIM", "IGG", "JAY"}

	data := []byte("11223344AMYBOBCAMDANEDDFAEGUYHIMIGGJAY")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if testVal.TestVal != expectedVal {
		t.Error("Unexpected results.")
		t.Errorf("Unexpected results.\nExpected:%v\nResult:%v\n", expectedVal, testVal.TestVal)
		t.Fail()
	}

	if testVal.Names != expectedNames {
		t.Error("Unexpected results.")
		t.Errorf("Unexpected results.\nExpected:%v\nResult:%v\n", expectedNames, testVal.Names)
		t.Fail()

	}
}

func TestArrayNestedStructParse(t *testing.T) {
	type Name struct {
		NameData string `flatfile:"2,2"`
	}
	type FfpTest struct {
		Names [3]Name `flatfile:"1,3"`
	}

	testVal := &FfpTest{}

	expectedNames := [3]Name{Name{NameData: "MY"}, Name{NameData: "OB"}, Name{NameData: "AM"}}

	data := []byte("AMYBOBCAM")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if testVal.Names != expectedNames {
		t.Error("Unexpected results.")
		t.Errorf("Unexpected results.\nExpected:%v\nResult:%v\n", expectedNames, testVal.Names)
		t.Fail()

	}
}

func TestSliceParse(t *testing.T) {
	type FfpTest struct {
		TestVal []int    `flatfile:"1,2,4"`
		Names   []string `flatfile:"9,3,10"`
	}

	testVal := &FfpTest{}
	expectedVal := []int{11, 22, 33, 44}

	expectedNames := []string{"AMY", "BOB", "CAM", "DAN", "EDD", "FAE", "GUY", "HIM", "IGG", "JAY"}

	data := []byte("11223344AMYBOBCAMDANEDDFAEGUYHIMIGGJAY")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for i := 0; i < len(testVal.TestVal); i++ {
		if testVal.TestVal[i] != expectedVal[i] {
			t.Error("Unexpected results.")
			t.Errorf("Unexpected results.\nExpected:%v\nResult:%v\n", expectedVal, testVal.TestVal)
			t.Fail()
		}
	}

	for i := 0; i < len(testVal.Names); i++ {
		if testVal.Names[i] != expectedNames[i] {
			t.Error("Unexpected results.")
			t.Errorf("Unexpected results.\nExpected:%v\nResult:%v\n", expectedNames, testVal.Names)
			t.Fail()
		}
	}
}

func TestSliceNestedStructParse(t *testing.T) {
	type Name struct {
		NameData string `flatfile:"2,2"`
	}
	type FfpTest struct {
		Names []Name `flatfile:"1,3,3"`
	}

	testVal := &FfpTest{}

	expectedNames := [3]Name{Name{NameData: "MY"}, Name{NameData: "OB"}, Name{NameData: "AM"}}

	data := []byte("AMYBOBCAM")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for i := 0; i < len(testVal.Names); i++ {
		if testVal.Names[i] != expectedNames[i] {
			t.Error("Unexpected results.")
			t.Errorf("Unexpected results.\nExpected:%v\nResult:%v\n", expectedNames, testVal.Names)
			t.Fail()
		}
	}
}

func TestOffsetParse(t *testing.T) {
	type Name struct {
		NameData     string `flatfile:"1,3"`
		Age          int    `flatfile:"4,3"`
		CurrencyPref string `flatfile:"7,3"`
	}

	testVal := &Name{}

	expectedName := &Name{NameData: "AMY", Age: int(123), CurrencyPref: "CAD"}

	data := [][]byte{[]byte("AMY"), []byte("123"), []byte("CAD")}

	for idx, dataval := range data {
		err := Unmarshal(dataval, testVal, idx, 0, true)
		t.Log(testVal)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}

	if testVal.NameData != expectedName.NameData || testVal.Age != expectedName.Age || testVal.CurrencyPref != expectedName.CurrencyPref {
		t.Error("Unexpected results.")
		t.Errorf("Unexpected results.\nExpected:%v\nResult:%v\n", expectedName, testVal)
		t.Fail()
	}
}

func TestCalcNumFieldsToUnmarshal(t *testing.T) {
	type Profile struct {
		NameData string `flatfile:"1,9"`
		Age      int    `flatfile:"10,2"`
	}

	var tests = []struct {
		MyProfile   *Profile
		Record      []byte
		IndexOffset int
		Want        int
	}{
		{&Profile{}, []byte("1234567891"), 0, 1},
		{&Profile{}, []byte("12345678910"), 0, 2},
		{&Profile{}, []byte("1234567"), 0, 0},
		{&Profile{}, []byte("1"), 1, 0},
		{&Profile{}, []byte("12"), 1, 1},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("CalcNumFieldsToUnmarshal-%d", idx)
		t.Run(testName, func(t *testing.T) {
			got, _, err := CalcNumFieldsToUnmarshal(tt.Record, tt.MyProfile, tt.IndexOffset)
			if err != nil {
				t.Errorf("CalcNumFieldsToUnmarshal(%v,%v,%d) got: %d want: %d err: %s", tt.Record, tt.MyProfile, tt.IndexOffset, got, tt.Want, err)
			}
			if got != tt.Want {
				t.Errorf("CalcNumFieldsToUnmarshal(%v,%v,%d) got: %d want: %d", tt.Record, tt.MyProfile, tt.IndexOffset, got, tt.Want)
			}
		})
	}
}
func TestCalcNumFieldsToUnmarshalRemainder(t *testing.T) {
	type Profile struct {
		NameData string `flatfile:"1,9"`
		Age      int    `flatfile:"10,2"`
	}

	var tests = []struct {
		MyProfile   *Profile
		Record      []byte
		IndexOffset int
		Want        []byte
	}{
		{&Profile{}, []byte("1234567891"), 0, []byte("1")},
		{&Profile{}, []byte("12345678910"), 0, []byte("")},
		{&Profile{}, []byte("1234567"), 0, []byte("1234567")},
		{&Profile{}, []byte("1"), 1, []byte("1")},
		{&Profile{}, []byte("12"), 1, []byte("")},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("CalcNumFieldsToUnmarshalRemainder-%d", idx)
		t.Run(testName, func(t *testing.T) {
			_, got, err := CalcNumFieldsToUnmarshal(tt.Record, tt.MyProfile, tt.IndexOffset)
			if err != nil {
				t.Errorf("err: %s", err)
			}
			if !bytes.Equal(got, tt.Want) {
				t.Errorf("CalcNumFieldsToUnmarshalRemainder(%s,%v,%d) got: %s want: %s", string(tt.Record), tt.MyProfile, tt.IndexOffset, string(got), string(tt.Want))
			}
		})
	}
}

func TestByte_Unmarshal(t *testing.T) {
	type ByteStruct struct {
		ByteOne byte `flatfile:"1,1,override=byte"`
		ByteTwo byte `flatfile:"2,1,override=byte"`
	}

	var tests = []struct {
		ByteSt ByteStruct
		Record []byte
		Want   ByteStruct
	}{
		{ByteStruct{}, []byte("1134567891"), ByteStruct{byte('1'), byte('1')}},
		{ByteStruct{}, []byte("a1345678910"), ByteStruct{'a', '1'}},
		{ByteStruct{}, []byte("a2"), ByteStruct{'a', '2'}},
		{ByteStruct{}, []byte("/a"), ByteStruct{'/', 'a'}},
		{ByteStruct{}, []byte("?1"), ByteStruct{'?', '1'}},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("TestByte_Unmarshal-%d", idx)
		t.Run(testName, func(t *testing.T) {
			err := Unmarshal(tt.Record, &tt.ByteSt, 0, 0, false)
			if err != nil {
				t.Errorf("err: %s", err)
			}
			if tt.ByteSt != tt.Want {
				t.Errorf("Unmarshal(%s,%v,0,0) got: %v want: %v", string(tt.Record), tt.ByteSt, tt.ByteSt, tt.Want)
			}
		})
	}
}

func TestRune_Unmarshal(t *testing.T) {
	type RuneStruct struct {
		RuneOne rune `flatfile:"1,1,override=rune"`
		RuneTwo rune `flatfile:"2,1,override=rune"`
	}

	var tests = []struct {
		RuneSt RuneStruct
		Record []byte
		Want   RuneStruct
	}{
		{RuneStruct{}, []byte("1134567891"), RuneStruct{'1', '1'}},
		{RuneStruct{}, []byte("a1345678910"), RuneStruct{'a', '1'}},
		{RuneStruct{}, []byte("a2"), RuneStruct{'a', '2'}},
		{RuneStruct{}, []byte("/a"), RuneStruct{'/', 'a'}},
		{RuneStruct{}, []byte("?1"), RuneStruct{'?', '1'}},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("TestRune_Unmarshal-%d", idx)
		t.Run(testName, func(t *testing.T) {
			err := Unmarshal(tt.Record, &tt.RuneSt, 0, 0, false)
			if err != nil {
				t.Errorf("err: %s", err)
			}
			if tt.RuneSt != tt.Want {
				t.Errorf("Unmarshal(%s,%v,0,0) got: %v want: %v", string(tt.Record), tt.RuneSt, tt.RuneSt, tt.Want)
			} else {
				t.Logf("Unmarshal(%s,%v,0,0) got: %v want: %v", string(tt.Record), tt.RuneSt, tt.RuneSt, tt.Want)
			}
		})
	}
}

func TestStartFieldIdx_Unmarshal(t *testing.T) {
	type ByteStruct struct {
		ByteOne byte `flatfile:"1,1,override=byte"`
		ByteTwo byte `flatfile:"2,1,override=byte"`
	}

	var tests = []struct {
		ByteSt ByteStruct
		Record []byte
		Want   ByteStruct
	}{
		{ByteStruct{}, []byte("1134567891"), ByteStruct{ByteTwo: byte('1')}},
		{ByteStruct{}, []byte("a1345678910"), ByteStruct{ByteTwo: byte('1')}},
		{ByteStruct{}, []byte("a2"), ByteStruct{ByteTwo: byte('2')}},
		{ByteStruct{}, []byte("/a"), ByteStruct{ByteTwo: byte('a')}},
		{ByteStruct{}, []byte("?1"), ByteStruct{ByteTwo: byte('1')}},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("TestByte_Unmarshal-%d", idx)
		t.Run(testName, func(t *testing.T) {
			err := Unmarshal(tt.Record, &tt.ByteSt, 1, 0, false)
			if err != nil {
				t.Errorf("err: %s", err)
			}
			if tt.ByteSt != tt.Want {
				t.Errorf("Unmarshal(%s,%v,0,0) data: %v got: %v want: %v", string(tt.Record), tt.Record, tt.ByteSt, tt.ByteSt, tt.Want)
			}
		})
	}
}

func TestNotAPointerErr_Umarshal(t *testing.T) {

	type FfpTest struct {
		TestVal string `flatfile:"1,1"`
	}

	testVal := FfpTest{}
	data := []byte("2")

	err := Unmarshal(data, testVal, 0, 0, false)

	if err == nil {
		t.Error("Unmarshal should return not a pointer error when attempting to unmarshal non-pointer type")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestNestedPointer_Unmarshal(t *testing.T) {
	type NestedStruct struct {
		ByteOne *int `flatfile:"1,1"`
	}

	type ByteStruct struct {
		NestedData *NestedStruct `flatfile:"1,1"`
	}

	var wantByte int
	wantByte = int(1)

	initByte := int(0)
	var GotByte *int
	GotByte = &initByte
	var tests = []struct {
		ByteSt ByteStruct
		Record []byte
		Want   ByteStruct
	}{
		{ByteStruct{NestedData: &NestedStruct{ByteOne: GotByte}}, []byte("1134567891"), ByteStruct{NestedData: &NestedStruct{ByteOne: &wantByte}}},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("TestByte_Unmarshal-%d", idx)
		t.Run(testName, func(t *testing.T) {
			err := Unmarshal(tt.Record, &tt.ByteSt, 0, 0, false)
			if err != nil {
				t.Errorf("err: %s", err)
			}
			if *tt.ByteSt.NestedData.ByteOne != *tt.Want.NestedData.ByteOne {
				t.Errorf("Unmarshal(%s,0,0,false) data: %v got: %d want: %d", string(tt.Record), tt.Record, tt.ByteSt.NestedData.ByteOne, tt.Want.NestedData.ByteOne)
			}
		})
	}
}

func TestShouldUnmarshal(t *testing.T) {

	var tests = []struct {
		ffTag  flatfileTag
		Record []byte
		Want   bool
	}{
		{flatfileTag{condCol: 1, condLen: 1, condVal: "1", condChk: true}, []byte("1134567891"), true},
		{flatfileTag{condCol: 1, condLen: 1, condVal: "9", condChk: true}, []byte("1134567891"), false},
		{flatfileTag{condCol: 4, condLen: 2, condVal: "45", condChk: true}, []byte("1134567891"), true},
		{flatfileTag{condCol: 4, condLen: 7, condVal: "456789", condChk: true}, []byte("1134567891"), false},
		{flatfileTag{condCol: 1, condLen: 1, condVal: "1", condChk: false}, []byte("1134567891"), true},
		{flatfileTag{condCol: 1, condLen: 1, condVal: "9", condChk: false}, []byte("1134567891"), true},
		{flatfileTag{condCol: 4, condLen: 2, condVal: "45", condChk: false}, []byte("1134567891"), true},
		{flatfileTag{condCol: 4, condLen: 7, condVal: "456789", condChk: false}, []byte("1134567891"), true},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("TestByte_Unmarshal-%d", idx)
		t.Run(testName, func(t *testing.T) {
			result := ShouldUnmarshal(&tt.ffTag, tt.Record)
			if result != tt.Want {
				t.Errorf("ShouldUnmarshal(%#v, %s) got: %v want: %v", tt.ffTag, string(tt.Record), result, tt.Want)
			}
		})
	}
}
