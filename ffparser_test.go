package ffparser

import (
	"reflect"
	"testing"
)

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

	err := Unmarshal(data, testVal)

	if err != nil {
		t.Error(err)
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

	err := Unmarshal(data, testVal)

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

	err := Unmarshal(data, testVal)

	if err == nil {
		t.Error("Unmarshal should return error when failing to parse bool")
	}
	t.Log(err)
}
