package main

import "fmt"

var locations = []string{
	"Atlanta",
	"Portland",
	"New York",
}

func main() {
	for _, name := range locations {
		fmt.Printf("Name: %s\n", name)
	}
	middle := locations[1:2]
	fmt.Printf("%s\n", middle)
}
