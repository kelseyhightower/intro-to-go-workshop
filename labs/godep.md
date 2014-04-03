# Managing Dependencies with Godep

## Install Godep

    go get github.com/tools/godep

## Inspect the dependencies for csv2json-cli

    cd ${GOPATH}/src/github.com/${username}/csv2json-cli
    go list -json

## Vendor Dependencies

    godep save

### Inspect Godeps/Godeps.json 

	{
		"ImportPath": "github.com/kelseyhightower/csv2json-cli",
		"GoVersion": "go1.2",
		"Deps": [
			{
				"ImportPath": "github.com/kelseyhightower/csv2json",
				"Rev": "0ca8ee22d850f992f466c21bb8105ef8cba2184b"
			}
		]
	}

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
