package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	infile string
)

type famousGopher struct {
	Name  string
	Date  string
	Title string
}

func init() {
	flag.StringVar(&infile, "infile", "", "input file")
}

func csv2json(r io.Reader) ([]byte, error) {
	gophers := make([]famousGopher, 0)

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

		gopher := famousGopher{
			Name:  record[0],
			Date:  record[1],
			Title: record[2],
		}
		gophers = append(gophers, gopher)
	}

	data, err := json.MarshalIndent(&gophers, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
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

	jsonData, err := csv2json(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
