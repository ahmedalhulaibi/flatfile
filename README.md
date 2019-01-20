# ffparser

This library provides a method `Unmarshal` which maps a slice of bytes to a struct based on defined struct tags.

Struct tags are in the form `ffp:"pos,len"`

Example: 
```go
type FileHeader struct {
	LogicalRecordTypeID byte   `ffp:"1,1"`
	LogicalRecordCount  uint32 `ffp:"2,9"`
	OriginatorID        string `ffp:"11,10"`
}
```
TODO:
- [ ] Marshal method
- [ ] Flat File abstraction