package ffparser

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

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

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int16")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt32(t *testing.T) {
	type Int32Struct struct {
		Int32One int32 `ffp:"1,11"`
		Int32Two int32 `ffp:"12,10"`
	}

	testVal := &Int32Struct{}
	data := []byte("-21474836482147483647")

	err := Unmarshal(data, testVal, 0, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int32One != -2147483648 || testVal.Int32Two != 2147483647 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt32InvalidSyntaxErr(t *testing.T) {
	type Int32Struct struct {
		Int32One int32 `ffp:"1,1"`
	}

	testVal := &Int32Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt32OutOfRangeErr(t *testing.T) {
	type Int32Struct struct {
		Int32One int32 `ffp:"1,10"`
	}

	testVal := &Int32Struct{}
	data := []byte("9999999999")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt64(t *testing.T) {
	type Int64Struct struct {
		Int64One int64 `ffp:"1,20"`
		Int64Two int64 `ffp:"21,19"`
	}

	testVal := &Int64Struct{}
	data := []byte("-92233720368547758089223372036854775807")

	err := Unmarshal(data, testVal, 0, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Int64One != -9223372036854775808 || testVal.Int64Two != 9223372036854775807 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestInt64InvalidSyntaxErr(t *testing.T) {
	type Int64Struct struct {
		Int64One int64 `ffp:"1,1"`
	}

	testVal := &Int64Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestInt64OutOfRangeErr(t *testing.T) {
	type Int64Struct struct {
		Int64One int64 `ffp:"1,19"`
	}

	testVal := &Int64Struct{}
	data := []byte("9999999999999999999")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse int64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFloat32(t *testing.T) {
	type Float32Struct struct {
		Float32One float32 `ffp:"1,22"`
		Float32Two float32 `ffp:"23,21"`
	}

	testVal := &Float32Struct{}
	data := []byte("3.4028234663852886e+381.401298464324817e-45")

	err := Unmarshal(data, testVal, 0, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Float32One != math.MaxFloat32 || testVal.Float32Two != math.SmallestNonzeroFloat32 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestFloat32InvalidSyntaxErr(t *testing.T) {
	type Float32Struct struct {
		Float32One float32 `ffp:"1,1"`
	}

	testVal := &Float32Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse Float32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFloat32OutOfRangeErr(t *testing.T) {
	type Float32Struct struct {
		Float32One float32 `ffp:"1,40"`
	}

	testVal := &Float32Struct{}
	data := []byte("9999999999999999999999999999999999999999")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse Float32")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFloat64(t *testing.T) {
	type Float64Struct struct {
		Float64One float64 `ffp:"1,23"`
		Float64Two float64 `ffp:"24,6"`
	}

	testVal := &Float64Struct{}
	data := []byte("1.7976931348623157e+3085e-324")

	err := Unmarshal(data, testVal, 0, 0)

	if err != nil {
		t.Error(err)
	}

	if testVal.Float64One != math.MaxFloat64 || testVal.Float64Two != math.SmallestNonzeroFloat64 {
		t.Log(testVal)
		t.Fail()
	}
}

func TestFloat64InvalidSyntaxErr(t *testing.T) {
	type Float64Struct struct {
		Float64One float64 `ffp:"1,1"`
	}

	testVal := &Float64Struct{}
	data := []byte("$")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse Float64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFloat64OutOfRangeErr(t *testing.T) {
	type Float64Struct struct {
		Float64One float64 `ffp:"1,23"`
	}

	testVal := &Float64Struct{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse Float64")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParsePosSyntaxErr(t *testing.T) {
	type FfpTest struct {
		TestVal string `ffp:"asdf,1"`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return syntax error when failing to parse position param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseLenSyntaxErr(t *testing.T) {
	type FfpTest struct {
		TestVal string `ffp:"1,asdf"`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return syntax error when failing to parse length param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParsePosRangeErr(t *testing.T) {
	type FfpTest struct {
		TestVal string `ffp:"-1,10"`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return out of range error when failing to parse position param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseLenRangeErr(t *testing.T) {
	type FfpTest struct {
		TestVal string `ffp:"1,-1"`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return out of range error when failing to parse length param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseMissingParamErr(t *testing.T) {
	type FfpTest struct {
		TestVal string `ffp:""`
	}

	testVal := &FfpTest{}
	data := []byte("2.7976931348623157e+308")

	err := Unmarshal(data, testVal, 0, 0)

	if err == nil {
		t.Error("Unmarshal should return missing parameter error")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestArrayParse(t *testing.T) {
	type FfpTest struct {
		TestVal [4]int     `ffp:"1,2"`
		Names   [10]string `ffp:"9,3"`
	}

	testVal := &FfpTest{}
	expectedVal := [4]int{11, 22, 33, 44}

	expectedNames := [10]string{"AMY", "BOB", "CAM", "DAN", "EDD", "FAE", "GUY", "HIM", "IGG", "JAY"}

	data := []byte("11223344AMYBOBCAMDANEDDFAEGUYHIMIGGJAY")

	err := Unmarshal(data, testVal, 0, 0)

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
		NameData string `ffp:"1,3"`
	}
	type FfpTest struct {
		Names [3]Name `ffp:"1,3"`
	}

	testVal := &FfpTest{}

	expectedNames := [3]Name{Name{NameData: "AMY"}, Name{NameData: "BOB"}, Name{NameData: "CAM"}}

	data := []byte("AMYBOBCAM")

	err := Unmarshal(data, testVal, 0, 0)

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
		TestVal []int    `ffp:"1,2,4"`
		Names   []string `ffp:"9,3,10"`
	}

	testVal := &FfpTest{}
	expectedVal := []int{11, 22, 33, 44}

	expectedNames := []string{"AMY", "BOB", "CAM", "DAN", "EDD", "FAE", "GUY", "HIM", "IGG", "JAY"}

	data := []byte("11223344AMYBOBCAMDANEDDFAEGUYHIMIGGJAY")

	err := Unmarshal(data, testVal, 0, 0)

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
		NameData string `ffp:"1,3"`
	}
	type FfpTest struct {
		Names []Name `ffp:"1,3,3"`
	}

	testVal := &FfpTest{}

	expectedNames := [3]Name{Name{NameData: "AMY"}, Name{NameData: "BOB"}, Name{NameData: "CAM"}}

	data := []byte("AMYBOBCAM")

	err := Unmarshal(data, testVal, 0, 0)

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
		NameData     string `ffp:"1,3"`
		Age          int    `ffp:"4,3"`
		CurrencyPref string `ffp:"7,3"`
	}

	testVal := &Name{}

	expectedName := &Name{NameData: "AMY", Age: int(123), CurrencyPref: "CAD"}

	data := [][]byte{[]byte("AMY"), []byte("123"), []byte("CAD")}

	for idx, dataval := range data {
		err := Unmarshal(dataval, testVal, idx, 0)
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

func TestCalcNumFieldsToMarshal(t *testing.T) {
	type Profile struct {
		NameData string `ffp:"1,9"`
		Age      int    `ffp:"10,2"`
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
		testName := fmt.Sprintf("CalcNumFieldsToMarshal-%d", idx)
		t.Run(testName, func(t *testing.T) {
			got, _, err := CalcNumFieldsToMarshal(tt.Record, tt.MyProfile, tt.IndexOffset)
			if err != nil {
				t.Errorf("CalcNumFieldsToMarshal(%v,%v,%d) got: %d want: %d err: %s", tt.Record, tt.MyProfile, tt.IndexOffset, got, tt.Want, err)
			}
			if got != tt.Want {
				t.Errorf("CalcNumFieldsToMarshal(%v,%v,%d) got: %d want: %d", tt.Record, tt.MyProfile, tt.IndexOffset, got, tt.Want)
			}
		})
	}
}
func TestCalcNumFieldsToMarshalRemainder(t *testing.T) {
	type Profile struct {
		NameData string `ffp:"1,9"`
		Age      int    `ffp:"10,2"`
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
		testName := fmt.Sprintf("CalcNumFieldsToMarshalRemainder-%d", idx)
		t.Run(testName, func(t *testing.T) {
			_, got, err := CalcNumFieldsToMarshal(tt.Record, tt.MyProfile, tt.IndexOffset)
			if err != nil {
				t.Errorf("err: %s", err)
			}
			if !bytes.Equal(got, tt.Want) {
				t.Errorf("CalcNumFieldsToMarshalRemainder(%s,%v,%d) got: %s want: %s", string(tt.Record), tt.MyProfile, tt.IndexOffset, string(got), string(tt.Want))
			}
		})
	}
}
