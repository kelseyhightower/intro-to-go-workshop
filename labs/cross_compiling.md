# Cross Compiling

### Requirements

- gcc

### Build the Tool Chains

#### Change Directory

    cd /usr/local/go/src

#### Run

    for os in linux windows darwin; do GOOS=${os} GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean; done


### Cross Compile csv2json_server

#### Change Directory

    cd ${GOPATH}/src/github.com/${username}/csv2json-server

#### Run

    GOOS=darwin GOARCH=amd64 go build -o csv2json-server main.go
