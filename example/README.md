# Usage

To run any example, navigate to the example directory and `go run main.go`. Some examples have specific instructions. See Below.

# [Basic Read File](basicReadFile/)

This example demonstrates a straightforward use case of reading a line from a file and unmarshalling it to a struct.

`go run main.go`

# [Buffered Read File](bufferedReadFile/)

This example demonstrates a use case of unmarshalling data when using buffered read and the line read exceeds the buffer size. 

2 options are shown.

1. Append the buffered data until the entire line is read into memory, then unmarshal.

`go run main.go 1`

2. Unmarshal the data as it is read.

`go run main.go 2`

# [Marshal Slice of Bytes to Struct](marshalBytes/)

This example demonstrates unmarshalling data from any slice of bytes `[]byte`. 

`go run main.go`


# [Tag named options](namedOptions/)

This example demonstrates the named tag options.

`go run main.go`


# [Override for byte and rune type](override/)

This example demonstrates the override option used to unmarshl byte and rune data types. 

`go run main.go`

This is required for byte and rune to be unmarshalled correctly as these types are aliases. The reflect package does not identify byte and rune as types and instead the types are returned as uint8 and int32.

# [Read into array](readIntoArray/)

This example demonstrates how data is unmarshalled into an array. If there is a data of the same typ that repeats in a single line of text

`go run main.go`

# [Read into slice by specifying occurrences](readIntoSlice/)

This example demonstrates how data is unmarshalled into a slice using the `occurs` tag option.

`go run main.go`