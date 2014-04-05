# Cross Compiling

### Requirements

- gcc

### Build the Linux Tool Chain

    cd /usr/local/go/src

#### Run

    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean


### Build the OS X Tool Chain

#### Run

    GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean


### Build the Windows Tool Chain

#### Run

    GOOS=windows GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean


### Cross Compile csv2json_server

-

    cd ${GOPATH}/src/github.com/${username}/csv2json-server

#### Run

    GOOS=darwin GOARCH=amd64 go build -o csv2json-server main.go
