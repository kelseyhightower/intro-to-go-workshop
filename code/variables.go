package main

import "fmt"

var (
	name     string
	Location = "Portland"
)

func main() {
	name = "Kelsey Hightower"
	distro := "CoreOS"
	fmt.Printf("Name: %s\nLocation: %s\nDistro: %s\n", name, Location, distro)
}
