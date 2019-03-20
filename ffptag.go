package ffparser

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ffpTagType struct {
	col      int
	length   int
	occurs   int
	override string
}

var parseFuncMap = map[string]func(string, *ffpTagType) error{
	"col":      parseColumnOption,
	"column":   parseColumnOption,
	"len":      parseLengthOption,
	"length":   parseLengthOption,
	"occ":      parseOccursOption,
	"occurs":   parseOccursOption,
	"ovr":      parseOverrideOption,
	"override": parseOverrideOption,
}

//validOptions is a list of the keys loaded from parseFuncMap. This is used purely to display options to user
var validOptions []string

func init() {
	for key := range parseFuncMap {
		validOptions = append(validOptions, key)
	}
}

//parseFfpTag parses an ffp struct tag on a field
//Tags are expected to be in the form:
// col,len,occurs
// where col is an int > 0
//		 len is an int
func parseFfpTag(fieldTag string, ffpTag *ffpTagType) error {
	var err error
	//split tag by comma to get column and length data
	params := strings.Split(fieldTag, ",")
	//column and length parameters must be provided
	if len(params) < 2 {
		return errors.Errorf("ffparser.parseFfpTag: Not enough ffp tag params provided.\nColumn and length parameters must be provided.\nMust be in form `ffp:\"col,len\"`")
	}

	for idx, param := range params {
		//check whether or not tag is using named options
		if strings.Contains(param, "=") {
			options := strings.Split(param, "=")
			if len(options) < 2 {
				return errors.Errorf("ffparser.parseFfpTag: Invalid formatting of named option '%v'\nNamed options should be in the form option=value\nValid options:%v", options, validOptions)
			}
			if funcVal, exists := parseFuncMap[options[0]]; exists {
				err = funcVal(options[1], ffpTag)
			} else {
				return errors.Errorf("ffparser.parseFfpTag: Invalid tag parameter %s\nValid options: %v", options[0], validOptions)
			}
		} else {
			//assume user is using positional options
			switch idx {
			case 0:
				err = parseColumnOption(param, ffpTag)
			case 1:
				err = parseLengthOption(param, ffpTag)
			case 2:
				err = parseOccursOption(param, ffpTag)
			case 3:
				err = parseOverrideOption(param, ffpTag)
			}
		}
		if err != nil {
			return errors.Wrapf(err, "ffparser.parseFfpTag: Error parsing tag option %s", param)
		}
	}

	if ffpTag.length == 0 || ffpTag.col == 0 {
		return errors.New("ffparser.parseFfpTag: Column or length option not provided")
	}
	return nil
}

func parseColumnOption(param string, ffpTag *ffpTagType) error {
	col, colerr := strconv.Atoi(param)
	if colerr != nil {
		return errors.Wrapf(colerr, "ffparser.parseColumnOption: Error parsing tag column parameter %s", param)
	}

	if col < 1 {
		return errors.Errorf("ffparser.parseColumnOption: Out of range error. Column parameter cannot be less than 1. Please note column is 1-indexed not zero")
	}
	ffpTag.col = col
	return nil
}

func parseLengthOption(param string, ffpTag *ffpTagType) error {

	length, lenerr := strconv.Atoi(param)
	if lenerr != nil {
		return errors.Wrapf(lenerr, "ffparser.parseLengthOption: Error parsing tag length parameter %s", param)
	}

	if length < 1 {
		return errors.Errorf("ffparser.parseLengthOption: Out of range error. Length parameter cannot be less than 1")
	}

	ffpTag.length = length
	return nil
}

func parseOccursOption(param string, ffpTag *ffpTagType) error {
	occurs, occerr := strconv.Atoi(param)
	if occerr != nil {
		return errors.Wrapf(occerr, "ffparser.parseOccursOption: Error parsing tag occurs parameter %s", param)
	}

	if occurs < 2 {
		return errors.Errorf("ffparser.parseOccursOption: Out of range error. Occurs parameter cannot be less than 2")
	}

	ffpTag.occurs = occurs
	return nil
}

func parseOverrideOption(param string, ffpTag *ffpTagType) error {
	if isValidOverride(param) {
		ffpTag.override = param
		return nil
	}
	return errors.Errorf("ffparser.parseOverrideOption: Invalid override %s", param)
}

func isValidOverride(override string) bool {
	switch override {
	case "byte",
		"rune":
		return true
	}
	return false
}
