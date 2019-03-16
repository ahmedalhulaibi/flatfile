package ffparser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ffpTagType struct {
	col    int
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
		return fmt.Errorf("ffparser.parseFfpTag: Not enough ffp tag params provided.\nPosition and length parameters must be provided.\nMust be in form `ffp:\"pos,len\"`")
	}

	col, colerr := strconv.Atoi(params[0])
	if colerr != nil {
		return errors.Wrapf(colerr, "ffparser.parseFfpTag: Error parsing tag position parameter %s", params[0])
	}

	if col < 1 {
		return fmt.Errorf("ffparser.parseFfpTag: Out of range error. Position parameter cannot be less than 1. Please note position is 1-indexed not zero")
	}

	ffpTag.col = col

	length, lenerr := strconv.Atoi(params[1])
	if lenerr != nil {
		return errors.Wrapf(lenerr, "ffparser.parseFfpTag: Error parsing tag length parameter %s", params[1])
	}

	if length < 1 {
		return fmt.Errorf("ffparser.parseFfpTag: Out of range error. Length parameter cannot be less than 1")
	}

	ffpTag.length = length

	if len(params) > 2 {
		occurs, occerr := strconv.Atoi(params[2])
		if occerr != nil {
			return errors.Wrapf(occerr, "ffparser.parseFfpTag: Error parsing tag occurs parameter %s", params[2])
		}

		if occurs < 2 {
			return fmt.Errorf("ffparser.parseFfpTag: Out of range error. Occurs parameter cannot be less than 2")
		}

		ffpTag.occurs = occurs
	}

	return nil
}
