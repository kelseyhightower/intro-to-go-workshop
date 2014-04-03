# Managing Dependencies with Godep

## Install Godep

    go get github.com/tools/godep

## Inspect the dependencies for csv2json-cli

    cd ${GOPATH}/src/github.com/${username}/csv2json-cli
    go list -json

## Vendor Dependencies

    godep save

### Inspect Godep Directory

	Godeps/
	├── Godeps.json
	├── Readme
	└── _workspace
		└── src
			└── github.com
				└── kelseyhightower
					└── csv2json
						├── README.md
						├── csv2json.go
						└── csv2json_test.go

## Build Using Godep

    godep go build -o csv2json main.go

## Test Using Godep

    godep go test

## Restore Dependencies

    rm -rf ${GOPATH}/src/github.com/${username}/csv2json
    godep restore


## Updating Dependencies

    godep restore
    go get -u csv2json
    godep save
