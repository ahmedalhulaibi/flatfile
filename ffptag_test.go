package ffparser

import "testing"

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
