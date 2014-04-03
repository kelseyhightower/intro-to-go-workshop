# Hello World

#### Create

    mkdir -p ${GOPATH}/src/github.com/${username}/hello

#### Edit

    ${GOPATH}/src/github.com/${username}/hello/main.go

#### Code

	package main

	import (
		"fmt"
	)

	func main() {
		fmt.Println("Hello World")
	}

#### Build

    cd ${GOPATH}/src/github.com/${username}/hello
    go build -o hello main.go

#### Run

    ./hello
