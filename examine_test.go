package flatfile

import "testing"

func TestExamine(t *testing.T) {
	type NestedStruct struct {
		slice []byte
	}
	type ExamineStruct struct {
		data    NestedStruct `tagger:"test"`
		slice   []byte
		integer int
	}

	Examine([]ExamineStruct{})
}
