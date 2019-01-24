# ffparser

This library provides a method `Unmarshal` which maps a slice of bytes to a struct based on defined struct tags. 

Struct tags are in the form `ffp:"pos,len"`

The intent is to eliminate boilerplate code for reading data from a flat file and mapping it to the properties in a struct.

[Example](https://github.com/ahmedalhulaibi/ffparser/tree/master/example)

TODO:
- [ ] Flat File abstraction
- [ ] Support for conditional marshal
- [ ] Emulate [COBOL occurs clause](https://www.ibm.com/support/knowledgecenter/en/SS6SG3_4.2.0/com.ibm.entcobol.doc_4.2/PGandLR/tasks/tptbl03.htm)
