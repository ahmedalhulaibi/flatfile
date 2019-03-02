package ffparser

import (
	"fmt"
	"strconv"
	"strings"
)

type ffpTagType struct {
	pos    int
	length int
	occurs int
}

//parseFfpTag parses an ffp struct tag on a field
//Tags are expected to be in the form:
// pos,len,occurs
// where pos is an int > 0
//		 len is an int
func parseFfpTag(fieldTag string, ffpTag *ffpTagType) error {

	//split tag by comma to get position and length data
	params := strings.Split(fieldTag, ",")
	//position and length parameters must be provided
	//
	if len(params) < 2 {
		return fmt.Errorf("ffparser: Not enough ffp tag params provided.\nPosition and length parameters must be provided.\nMust be in form `ffp:\"pos,len\"`")
	}

	pos, poserr := strconv.Atoi(params[0])
	if poserr != nil {
		return fmt.Errorf("ffparser: Error parsing position parameter\n%s", poserr)
	}

	if pos < 1 {
		return fmt.Errorf("ffparser: Out of range error. Position parameter cannot be less than 1. Please note position is 1-indexed not zero")
	}

	ffpTag.pos = pos

	length, lenerr := strconv.Atoi(params[1])
	if lenerr != nil {
		return fmt.Errorf("ffparser: Error parsing length parameter\n%s", lenerr)
	}

	if length < 1 {
		return fmt.Errorf("ffparser: Out of range error. Length parameter cannot be less than 1")
	}

	ffpTag.length = length

	if len(params) > 2 {
		occurs, occerr := strconv.Atoi(params[2])
		if occerr != nil {
			return fmt.Errorf("ffparser: Error parsing occurs parameter\n%s", occerr)
		}

		if occurs < 2 {
			return fmt.Errorf("ffparser: Out of range error. Occurs parameter cannot be less than 2")
		}

		ffpTag.occurs = occurs
	}

	return nil
}
