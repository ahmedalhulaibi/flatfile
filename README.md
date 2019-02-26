# ffparser

This library provides a method `Unmarshal` which maps a slice of bytes to a struct based on defined struct tags. 

Struct tags are in the form `ffp:"pos,len"`

The intent is to eliminate boilerplate code for reading data from a flat file and mapping it to the properties in a struct.

[Example](https://github.com/ahmedalhulaibi/ffparser/tree/master/example)

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
- [ ] Slice
- [ ] Array
- [x] Nested struct

TODO:
- [ ] Slice, Array support AKA Emulate [COBOL occurs clause](https://www.ibm.com/support/knowledgecenter/en/SS6SG3_4.2.0/com.ibm.entcobol.doc_4.2/PGandLR/tasks/tptbl03.htm)
- [ ] Flat File abstraction
- [ ] Support for conditional marshal.
