package main

import (
	"fmt"
	"log"
)

func main() {
	m := make(map[string]string)
	m["name"] = "Kelsey"

	name, ok := m["name"]
	if !ok {
		log.Fatal("Name does not exist")
	}
	fmt.Println(name)

	for k, v := range m {
		fmt.Printf("Key: %s Value: %s", k, v)
	}
}
