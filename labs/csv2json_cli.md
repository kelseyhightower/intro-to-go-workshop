# csv2json-cli

CLI tool to convert CSV to JSON

#### Create

    ${GOPATH}/src/github.com/${username}/csv2json-cli

#### Edit

    ${GOPATH}/src/github.com/${username}/csv2json-cli/main.go

-

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

#### Edit

    ${HOME}/famous-gophers.csv

-

	Mac, 1947, The Goofy Gophers
	Tosh, 1947, The Goofy Gophers
	Samuel J. Gopher, 1966, Winnie the Pooh
	Chief Running Board, 1968, Go Go Gophers
	Ruffled Feathers, 1968, Go Go Gophers

#### Build

    go build -o csv2json .

#### Run

    ./csv2json -infile ${HOME}/famous-gophers.csv


## Exercise

### Generate Go docs

    godoc -http=":8000"
