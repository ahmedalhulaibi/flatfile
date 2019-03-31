package flatfile

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

type testType struct {
	Data   string `flatfile:"col=1,len=10"`
	Number int    `flatfile:"col=11,len=1"`
}

func TestFlatFileNew(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	_, err := New(reader, &testType{})
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestFlatFileNew_Err(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	_, err := New(reader, testType{})
	if err == nil {
		t.Errorf("Expected not a pointer error")
	}
	t.Log(err)
}

func TestFlatFileRead(t *testing.T) {

	testCases := []struct {
		desc     string
		lineData string
		got      *testType
		want     testType
		wantErr  error
	}{
		{
			desc:     "Read nothing",
			lineData: "",
			got:      &testType{},
			want:     testType{Data: "", Number: 0},
			wantErr:  io.EOF,
		},
		{
			desc:     "Read some data",
			lineData: "DATA!DATA!9",
			got:      &testType{},
			want:     testType{Data: "DATA!DATA!", Number: 9},
			wantErr:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tC.lineData))
			file, err := New(reader, tC.got)
			if err != nil {
				t.Errorf("Unexpected error %s", err.Error())
			}
			err = file.Read()
			if err != nil && err != tC.wantErr {
				t.Errorf("Unexpected error %s wantErr %v", err.Error(), tC.wantErr)
			}
			if tC.got.Data != tC.want.Data || tC.got.Number != tC.want.Number {
				t.Errorf("flatfile.Read() got %v want %v with error %v", tC.got, tC.want, tC.wantErr)
			}
		})
	}
}
