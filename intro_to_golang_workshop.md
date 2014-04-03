
## CLI tool to convert CSV to JSON - csv2json-cli

Convert CSV files into JSON.

### Write the code

Create `~/go/src/github.com/username/csv2json-cli`

Edit `~/go/src/github.com/username/csv2json-cli/main.go`

```
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/csv2json"
)

var (
	infile string
	columns = []string{"Name","Date","Title"}
)

func init() {
	flag.StringVar(&infile, "infile", "", "input file")
}

func main() {
	flag.Parse()
	if infile == "" {
		log.Fatal("Input file required")
	}
	f, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	jsonData, err := csv2json.Convert(f, columns)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
```

Edit ~/go/src/github.com/username/famous-gophers.csv

```
Mac, 1947, The Goofy Gophers
Tosh, 1947, The Goofy Gophers
Samuel J. Gopher, 1966, Winnie the Pooh
Chief Running Board, 1968, Go Go Gophers
Ruffled Feathers, 1968, Go Go Gophers
```

Build

```
go build main.go
```

Run

```
./csv2json-cli -infile famous-gophers.csv
```


### Managing dependencies with godep

```
go list -json
go get github.com/kr/godep
godep save
```

### Exercise: Better error messages

Improve the error message when err != nil

Hint 

```
import "error"
```

Exercise: Add comments

//

/* */


### Exercise: Generate Go docs

    godoc -http=":8080"

## HTTP API to convert CSV to JSON - csv2json-http

### Write the code

Create ~/go/src/github.com/username/csv2json-http

Edit ~/go/src/github.com/username/csv2json-http/main.go


```
package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/csv2json"
)

var (
	columns = []string{"Name", "Date", "Title"}
)

func csv2JsonServer(w http.ResponseWriter, req *http.Request) {
	jsonData, err := csv2json.Convert(req.Body, columns)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, "Could not convert csv to json", 503)
	}
	print(string(jsonData))
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/csv2json", csv2JsonServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```


Build

```
go build main.go
```


Run

```
./csv2json-http
```

### Exercise: Make the listen port configurable via the environment

Hint

```
import "os"
os.Getenv("PORT")
```

### Exercise: Add more logging

Hint

```
import "log"
```
