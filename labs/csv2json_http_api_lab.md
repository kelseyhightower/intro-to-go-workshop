# HTTP API to convert CSV to JSON - csv2json-http

## Code

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

## Build

```
go build main.go
```


## Run

```
./csv2json-http
```

## Exercises

### Make the listen port configurable via the environment

Hint

```
import "os"
os.Getenv("PORT")
```

### Add more logging

Hint

```
import "log"
```
