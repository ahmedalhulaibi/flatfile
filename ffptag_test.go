package flatfile

import (
	"fmt"
	"testing"
)

func TestFfpTagOptions_parseFfpTag(t *testing.T) {
	var tests = []struct {
		tagValue  string
		ResultTag *flatfileTag
		WantTag   *flatfileTag
		isError   bool
	}{
		{"1,1", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"1,1,2", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"1,1,2,byte", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "byte", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"1,1,2,rune", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "rune", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"1,1,2,byte,1-10-tenletters", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "byte", condChk: true, condCol: 1, condLen: 10, condVal: "tenletters"}, false},
		{"col=1,len=1", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"column=1,len=1", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"col=1,length=1", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"column=1,length=1", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"col=1,len=1,occ=2", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"col=1,len=1,occurs=2", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"col=1,len=1,occ=2,ovr=byte", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "byte", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"col=1,len=1,occ=2,override=rune", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "rune", condChk: false, condCol: 0, condLen: 0, condVal: ""}, false},
		{"col=1,len=1,occ=2,override=rune,cond=1-10-tenletters", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 2, override: "rune", condChk: true, condCol: 1, condLen: 10, condVal: "tenletters"}, false},
		{"override=rune,cond=3-1-1", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
		{"col=1=1,len=3", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
		{"1,2,fake=1", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
		{"col=1,len=3,occ=once", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
		{"col=1,len=3,occ=1", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
		{"col=1,len=3,occ=2,ovr=uintptr", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
		{"col=1,len=3,occ=2,ovr=byte,cond=1-2", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
		{"col=1,len=3,occ=2,ovr=byte,cond=one-2-to", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
		{"col=1,len=3,occ=2,ovr=byte,cond=1-two-to", &flatfileTag{}, &flatfileTag{col: 1, length: 1, occurs: 0, override: "", condChk: false, condCol: 0, condLen: 0, condVal: ""}, true},
	}

	for idx, tt := range tests {
		testName := fmt.Sprintf("TestFfpTagOptions_parseFfpTag-%d", idx)
		t.Run(testName, func(t *testing.T) {
			err := parseFlatfileTag(tt.tagValue, tt.ResultTag)
			if err != nil && !tt.isError {
				t.Errorf("parseFfpTag(%v,%v) got: %v want: %v err: %s", tt.tagValue, tt.ResultTag, tt.ResultTag, tt.WantTag, err)
			} else if err == nil && tt.isError {
				t.Errorf("parseFfpTag(%v,%v) Expected Error! got: %v want: %v err: %s", tt.tagValue, tt.ResultTag, tt.ResultTag, tt.WantTag, err)
			} else if err != nil && tt.isError {
				//success
				t.Logf("parseFfpTag(%v,%v) Success Error! got: %v want: %v err: %s", tt.tagValue, tt.ResultTag, tt.ResultTag, tt.WantTag, err)
			} else {
				if tt.ResultTag.col != tt.WantTag.col || tt.ResultTag.length != tt.WantTag.length || tt.ResultTag.occurs != tt.WantTag.occurs || tt.ResultTag.override != tt.WantTag.override || tt.ResultTag.condChk != tt.WantTag.condChk || tt.ResultTag.condCol != tt.WantTag.condCol || tt.ResultTag.condLen != tt.WantTag.condLen || tt.ResultTag.condVal != tt.WantTag.condVal {
					t.Errorf("parseFfpTag(%v,%v) got: %v want: %v", tt.tagValue, tt.ResultTag, tt.ResultTag, tt.WantTag)
				}
			}
		})
	}
}

func TestFfpTagParseMissingParamErr_parseFfpTag(t *testing.T) {
	testVal := `flatfile:""`

	err := parseFlatfileTag(testVal, &flatfileTag{})

	if err == nil {
		t.Error("parseFfpTag should return missing parameter error")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseLenRangeErr_parseFfpTag(t *testing.T) {
	testVal := `flatfile:"1,-1"`

	err := parseFlatfileTag(testVal, &flatfileTag{})

	if err == nil {
		t.Error("parseFfpTag should return out of range error when failing to parse length param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParsePosRangeErr_parseFfpTag(t *testing.T) {
	testVal := `flatfile:"-1,10"`

	err := parseFlatfileTag(testVal, &flatfileTag{})

	if err == nil {
		t.Error("parseFfpTag should return out of range error when failing to parse position param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParseLenSyntaxErr_parseFfpTag(t *testing.T) {
	testVal := `flatfile:"1,asdf"`

	err := parseFlatfileTag(testVal, &flatfileTag{})

	if err == nil {
		t.Error("parseFfpTag should return syntax error when failing to parse length param")
	}
	t.Log(testVal)
	t.Log(err)
}

func TestFfpTagParsePosSyntaxErr_parseFfpTag(t *testing.T) {
	testVal := `flatfile:"asdf,1"`

	err := parseFlatfileTag(testVal, &flatfileTag{})

	if err == nil {
		t.Error("parseFfpTag should return syntax error when failing to parse position param")
	}
	t.Log(testVal)
	t.Log(err)
}
