package flatfile

import (
	"bufio"
	"fmt"
	"io"
	"reflect"

	"github.com/pkg/errors"
)

//FlatFile is an abstraction for a flat file containing structured data
type FlatFile struct {
	reader       *bufio.Reader
	objectLayout interface{}
}

//New returns a new FlatFile reader object
func New(reader *bufio.Reader, objectLayout interface{}) (*FlatFile, error) {
	if reflect.TypeOf(objectLayout).Kind() == reflect.Ptr {
		return &FlatFile{reader: reader, objectLayout: objectLayout}, nil
	}

	return nil, errors.Wrap(fmt.Errorf("flatfile.New: %s is not a pointer", reflect.TypeOf(objectLayout)), "")
}

//Read will read a line from a bufio.Reader and call flatfile.Unmarshal to convert the read in data into FlatFile.objectLayout
func (f *FlatFile) Read() (err error) {
	var line []byte
	var buffLine []byte
	prefix := true
	for prefix {
		buffLine, prefix, err = f.reader.ReadLine()
		if err == io.EOF {
			return err
		}
		line = append(line, buffLine...)
	}

	return Unmarshal(line, f.objectLayout, 0, 0, false)
}
