# Hello World

#### Create

    mkdir ${GOPATH}/src/hello

#### Change Directory

    cd ${GOPATH}/src/hello

#### Edit

    main.go

-

	package main

	import (
		"fmt"
	)

	func main() {
		fmt.Println("Hello World")
	}

#### Build

    go build -o hello .

#### Run

    ./hello
