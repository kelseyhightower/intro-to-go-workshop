# csv2json-server

HTTP API to convert CSV to JSON

#### Create

    mkdir ${GOPATH}/src/github.com/${username}/csv2json-server

#### Change Directory

    cd ${GOPATH}/src/github.com/${username}/csv2json-server

#### Edit

    main.go

-

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
		w.Write(jsonData)
	}

	func main() {
		http.HandleFunc("/csv2json", csv2JsonServer)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}

#### Build

    go build -o csv2json-server .

#### Run

    ./csv2json-server

Test with curl:

    curl -X POST http://localhost:8080/csv2json --data-binary @${HOME}/famous-gophers.csv

#### Version

    git init .
    git add main.go
    git commit -m "first commit"

#### Vendor

    godep save
    git add Godeps
    git commit -m "Manage dependencies with godep"

## Exercise

### Improve logging

#### Hint

    import (
        "log"
    )
