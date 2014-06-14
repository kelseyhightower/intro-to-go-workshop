// START 0 OMIT
package main

import "fmt"

type Person struct {
	name     string
	Location string
}
// END 0 OMIT
// START 1 OMIT
func NewPerson(name string) *Person {
	p := Person{
		name: name,
	}
	return &p
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func main() {
	p := NewPerson("Kelsey")
	fmt.Printf("Name: %s\n", p.Name())
	p.Location = "Portland"
	fmt.Printf("Location: %s\n", p.Location)
}
// END 1 OMIT
