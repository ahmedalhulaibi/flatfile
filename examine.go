package ffparser

import (
	"fmt"
	"reflect"
	"strings"
)

// Examine traverses all elements of a type and uses the reflect pkg to print type and kind
func Examine(v interface{}) {
	examiner(reflect.TypeOf(v), 0)
}

// Below code is sourced from Jon Bodner's blog: https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
// Direct link to Gist: https://gist.github.com/jonbodner/1727d0825d73541db8d6fcb859515735
func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}
