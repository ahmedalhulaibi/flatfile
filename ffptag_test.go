package ffparser

import (
	"fmt"
	"testing"
)

func TestFfpTagPositionalOptions_parseFfpTag(t *testing.T) {
	var tests = []struct {
		tagValue  string
		ResultTag *ffpTagType
		WantTag   *ffpTagType
	}{
		{"1,1", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 0, override: ""}},
		{"1,1,2", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 2, override: ""}},
		{"1,1,2,byte", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 2, override: "byte"}},
		{"1,1,2,rune", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 2, override: "rune"}},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("TestFfpTagPositionalOptions_parseFfpTag-%d", idx)
		t.Run(testName, func(t *testing.T) {
			err := parseFfpTag(tt.tagValue, tt.ResultTag)
			if err != nil {
				t.Errorf("parseFfpTag(%v,%v) got: %v want: %v err: %s", tt.tagValue, tt.ResultTag, tt.ResultTag, tt.WantTag, err)
			}
			if tt.ResultTag.col != tt.WantTag.col || tt.ResultTag.length != tt.WantTag.length || tt.ResultTag.occurs != tt.WantTag.occurs || tt.ResultTag.override != tt.WantTag.override {
				t.Errorf("parseFfpTag(%v,%v) got: %v want: %v", tt.tagValue, tt.ResultTag, tt.ResultTag, tt.WantTag)
			}
		})
	}
}
func TestFfpTagNamedOptions_parseFfpTag(t *testing.T) {
	var tests = []struct {
		tagValue  string
		ResultTag *ffpTagType
		WantTag   *ffpTagType
	}{
		{"col=1,len=1", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 0, override: ""}},
		{"column=1,len=1", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 0, override: ""}},
		{"col=1,length=1", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 0, override: ""}},
		{"column=1,length=1", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 0, override: ""}},
		{"col=1,len=1,occ=2", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 2, override: ""}},
		{"col=1,len=1,occurs=2", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 2, override: ""}},
		{"col=1,len=1,occ=2,ovr=byte", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 2, override: "byte"}},
		{"col=1,len=1,occ=2,override=rune", &ffpTagType{}, &ffpTagType{col: 1, length: 1, occurs: 2, override: "rune"}},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("TestFfpTagPositionalOptions_parseFfpTag-%d", idx)
		t.Run(testName, func(t *testing.T) {
			err := parseFfpTag(tt.tagValue, tt.ResultTag)
			if err != nil {
				t.Errorf("parseFfpTag(%v,%v) got: %v want: %v err: %s", tt.tagValue, tt.ResultTag, tt.ResultTag, tt.WantTag, err)
			}
			if tt.ResultTag.col != tt.WantTag.col || tt.ResultTag.length != tt.WantTag.length || tt.ResultTag.occurs != tt.WantTag.occurs || tt.ResultTag.override != tt.WantTag.override {
				t.Errorf("parseFfpTag(%v,%v) got: %v want: %v", tt.tagValue, tt.ResultTag, tt.ResultTag, tt.WantTag)
			}
		})
	}
}
func TestFfpTagParseMissingParamErr_parseFfpTag(t *testing.T) {
	testVal := `ffp:""`

	err := parseFfpTag(testVal, &ffpTagType{})

	if err == nil {
		t.Error("parseFfpTag should return missing parameter error")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseLenRangeErr_parseFfpTag(t *testing.T) {
	testVal := `ffp:"1,-1"`

	err := parseFfpTag(testVal, &ffpTagType{})

	if err == nil {
		t.Error("parseFfpTag should return out of range error when failing to parse length param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParsePosRangeErr_parseFfpTag(t *testing.T) {
	testVal := `ffp:"-1,10"`

	err := parseFfpTag(testVal, &ffpTagType{})

	if err == nil {
		t.Error("parseFfpTag should return out of range error when failing to parse position param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseLenSyntaxErr_parseFfpTag(t *testing.T) {
	testVal := `ffp:"1,asdf"`

	err := parseFfpTag(testVal, &ffpTagType{})

	if err == nil {
		t.Error("parseFfpTag should return syntax error when failing to parse length param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParsePosSyntaxErr_parseFfpTag(t *testing.T) {
	testVal := `ffp:"asdf,1"`

	err := parseFfpTag(testVal, &ffpTagType{})

	if err == nil {
		t.Error("parseFfpTag should return syntax error when failing to parse position param")
	}
	t.Log(testVal)
	t.Log(err)
}
