# ffparser

The purpose of this library is provide a utility to read a record from structured [flat-file database](https://en.wikipedia.org/wiki/Flat-file_database).


This library provides a method `Unmarshal` which will convert a slice of bytes to fields in a struct based on defined struct tags. 

Struct tags are in the form `ffp:"pos,len"` or optionally `ffp:"pos,len,occurs"`

The intent is to eliminate boilerplate code for reading data from a flat file and mapping it to the properties in a struct.

Data type support:
- [x] bool
- [x] string
- [x] int
- [x] int8
- [x] int16
- [x] int32
- [x] int64
- [x] uint
- [x] uint8
- [x] uint16
- [x] uint32
- [x] uint64
- [x] float32
- [x] float64
- [x] Slice
- [x] Array
- [x] Nested struct

TODO:
- [x] Slice, Array support AKA Emulate [COBOL occurs clause](https://www.ibm.com/support/knowledgecenter/en/SS6SG3_4.2.0/com.ibm.entcobol.doc_4.2/PGandLR/tasks/tptbl03.htm)
- [ ] Flat File abstraction
- [ ] Support for conditional unmarshal 
      if field(pos,len) == "text" do unmarshal else skip
- [ ] Byte and Rune support using type override

# Usage

Use your favourite dependency tool to pull the code. dep, go mod, etc. Or use go get.

`go get github.com/ahmedalhulaibi/ffparser`

## [Examples](https://github.com/ahmedalhulaibi/ffparser/tree/master/example)
