# Cross Compiling

## Requirements

- gcc

## Build the Tool Chains

#### Change Directory

    cd /usr/local/go/src

#### Run

    for os in linux windows darwin; do GOOS=${os} GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean; done

**NOTE**: Make sure to not include the OS you are actually running in this list or you will end up with
a build of Go that has CGO disabled. So if you are running on OSX, for exampple, this should be:

    for os in linux windows; do GOOS=${os} GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean; done

## Cross Compile csv2json_server

#### Change Directory

    cd ${GOPATH}/src/github.com/${username}/csv2json-server

#### Run

    GOOS=darwin GOARCH=amd64 go build -o csv2json-server main.go

#### Inspect

    file csv2json-server

## Exercise

### Cross Compile for linux, windows, and darwin
