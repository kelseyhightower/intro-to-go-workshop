# Cross Compiling

### Key Environment Variables

- GOOS
- GOARCH
- CGO_ENABLED

### Requirements

- gcc

### Build the Linux Tool Chain

    cd /usr/local/go/src
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean

### Build the OS X Tool Chain

    cd /usr/local/go/src
    GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean

### Build the Windows Tool Chain

    cd /usr/local/go/src
    GOOS=windows GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean

### Cross Compile csv2json_server

    cd ${GOPATH}/src/github.com/${username}/csv2json-server
    GOOS=darwin GOARCH=amd64 go build -o csv2json-server main.go
