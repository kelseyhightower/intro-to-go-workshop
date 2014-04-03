# Package to convert CSV to JSON - csv2json

## Code

Create `~/go/src/github.com/username/csv2json`

Edit `~/go/src/github.com/username/csv2json/csv2json.go`

```
package csv2json

import (
	"encoding/csv"
	"encoding/json"
	"io"
)

func Convert(r io.Reader, columns []string) ([]byte, error) {
	rows := make([]map[string]string, 0)
	csvReader := csv.NewReader(r)
	csvReader.TrimLeadingSpace = true
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		row := make(map[string]string)
		for i, n := range columns {
			row[n] = record[i]
		}
		rows = append(rows, row)
	}
	data, err := json.MarshalIndent(&rows, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
```

## Build

```
go build main.go
```

## Test

Edit: `~/go/src/github.com/username/csv2json/csv2json_test.go`

```
package csv2json

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var columns = []string{"name", "email", "phone"}
var want = `[
  {
	"email": "ken@google.com",
	"name": "Ken Thompson",
	"phone": "555-5550"
  },
  {
	"email": "rob@google.com",
	"name": "Rob Pike",
	"phone": "555-5551"
  },
  {
	"email": "robert@google.com",
	"name": "Robert Griesemer",
	"phone": "555-5552"
  }
]`

var input = `
Ken Thompson, ken@google.com, 555-5550
Rob Pike, rob@google.com, 555-5551
Robert Griesemer, robert@google.com, 555-5552
`

func TestConvertWithBuffer(t *testing.T) {
	var b bytes.Buffer
	_, err := b.WriteString(input)
	if err != nil {
		t.Error(err)
	}
	got, err := Convert(bytes.NewReader(b.Bytes()), columns)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("TestConvertWithBuffer(f, %s) = %s; want %s",
			columns, string(got), want)
	}
}
```

### Running the test suite

```
go test
```

Run with the -v flag for verbose output:

```
go test -v
```

### Exercise: Write an additional test

Create an additional test named TestConvertWithFile.

**Hint**

```
func TestConvertWithFile(t *testing.T) {
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tf.Name())

	err = ioutil.WriteFile(tf.Name(), []byte(input), 0644)
	if err != nil {
		t.Error(err)
	}

	f, err := os.Open(tf.Name())
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	// Test your results
}
```
