# ffparser

The purpose of this package is provide a utility to read a record from a structured [flat-file database](https://en.wikipedia.org/wiki/Flat-file_database) or a record from a text file into a struct. The intent is to eliminate boilerplate code for reading data from a flat file and mapping it to the fields in a struct.

This is allows you to define your record layout mapping using struct tags.

This library provides a method `Unmarshal` which will convert a record (slice of bytes) into a struct.

# Usage

Use your favourite dependency tool to pull the code. dep, go mod, etc. Or use go get.

`go get github.com/ahmedalhulaibi/ffparser`

## Examples

Please refer to [examples](https://github.com/ahmedalhulaibi/ffparser/tree/master/example) folder in repo for all examples.

Let's say I have a text file `customers.txt` which contains customer data. There are 2 records:

```
AMY1900-01-01019123 FAKE STREETCA
BOB1800-01-01037456 OLD STREET US
```
The file layout is:
```
Columns  1, Length  3   = Customer Name
Columns  4, Length 10   = Customer Join/Open Date
Columns 14, Length  3   = Customer Age
Columns 17, Length 15   = Customer Street Address
Columns 32, Length  2   = Customer Country
```


Each field in a struct can be mapped to a single field in a record using a struct tag.

Struct tags are in the form `flatfile:"col,len"` or for a slice field `flatfile:"col,len,occurences"`.

I can directly translate these to fields in a struct:

```go
type CustomerRecord struct {
    //ffparser is one indexed, column starts at 1
    Name        string `flatfile:"1,3"`
    OpenDate    string `flatfile:"4,10"`
    Age         uint   `flatfile:"14,3"`
    Address     string `flatfile:"17,15"`
    CountryCode string `flatfile:"32,2"`
}

```

Once your layout is defined we can read a line from a file and unmarshal the data:

```go
type CustomerRecord struct {
	//ffparser is one indexed, column starts at 1
	Name        string `flatfile:"1,3"`
	OpenDate    string `flatfile:"4,10"`
	Age         uint   `flatfile:"14,3"`
	Address     string `flatfile:"17,15"`
	CountryCode string `flatfile:"32,2"`
}

func main() {
	file, err := os.Open("customers.txt")
	checkError(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	eof := false
	for !eof {
		data := readLine(reader)
		if data == nil {
			eof = true
		} else {
			fileRecord := &CustomerRecord{}
            //unmarhsal text data to struct
			err := ffparser.Unmarshal(data, fileRecord, 0, 0)
			fmt.Printf("%v\n", fileRecord)
			checkError(err)
		}
	}
	checkError(err)
}

func readLine(reader *bufio.Reader) []byte {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return nil
	}

	return str
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
```




# Features

## Data type support:
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
- [x] rune
- [x] byte
- [x] Slice
- [x] Array
- [x] Nested struct
- [x] Nested pointer (to any support type including struct)

- [x] Slice, Array support AKA Emulate [COBOL occurs clause](https://www.ibm.com/support/knowledgecenter/en/SS6SG3_4.2.0/com.ibm.entcobol.doc_4.2/PGandLR/tasks/tptbl03.htm)

- [x] Offset feature to support reading long lines of data. [Example](https://github.com/ahmedalhulaibi/ffparser/tree/master/example/bufferedReadFile)

- [x] Byte and Rune support using type override. 

    These are aliases for uint8 and int32 respectively. These require an override option to be supplied.
## TODO:
- [ ] Flat File abstraction
- [ ] Support for conditional unmarshal 
    
    if field(col,len) == "text" do unmarshal else skip. 
    
    This is useful for flat files where there are multiple record layouts within the same file.

