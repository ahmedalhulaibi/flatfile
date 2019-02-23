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

func TestUint8Err(t *testing.T) {
	type Uint8Struct struct {
		Uint8One uint8 `ffp:"1,1"`
		Uint8Two uint8 `ffp:"2,4"`
	}

	testVal := &Uint8Struct{}
	data := []byte("$255")

	err := Unmarshal(data, testVal, 0)

	if err == nil {
		t.Log(testVal)
		t.Error("Unmarshal should return error when failing to parse bool")
	}
	t.Log(err)
}
