# Cross Compiling

### Requirements

- gcc

### Build the Tool Chains


#### Run

    cd /usr/local/go/src
    for i in linux windows darwin; do GOOS=${i} GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean; done


### Cross Compile csv2json_server

#### Run

    cd ${GOPATH}/src/github.com/${username}/csv2json-server
    GOOS=darwin GOARCH=amd64 go build -o csv2json-server main.go
