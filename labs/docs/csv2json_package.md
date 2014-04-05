# csv2json

Package to convert CSV to JSON

#### Create

    mkdir ${GOPATH}/src/github.com/${username}/csv2json

#### Change Directory

    cd ${GOPATH}/src/github.com/${username}/csv2json

#### Edit

    csv2json.go

-

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

#### Build

    go install .

#### Inspect

    file ${GOPATH}/pkg/linux_amd64/github.com/${username}/csv2json.a

#### Version

    git init .
    git add .
    git commit -m "first commit"
