package main

import "fmt"

func main() {
	ints := make([]int, 0)
	for i := 0; i <= 100; i++ {
		ints = append(ints, i)
	}
	fmt.Printf("%#v", ints)
}
