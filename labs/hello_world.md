# Hello World

#### Create

    ${GOPATH}/src/hello

#### Edit

    ${GOPATH}/src/hello/main.go

-

	package main

	import (
		"fmt"
	)

	func main() {
		fmt.Println("Hello World")
	}

#### Build

    go build -o hello main.go

#### Run

    ./hello
