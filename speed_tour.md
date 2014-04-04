# Speed Tour

## Packages

### Creating a package

Create the package directory under GOPATH:

    mkdir ${GOPATH}/src/github.com/kelseyhightower/envconfig
    touch ${GOPATH}/src/github.com/kelseyhightower/envconfig/envconfig.go

Declare a package

    package envconfig

### Using packages

    import "github.com/kelseyhightower/envconfig"

Multiple imports

    import(
       "fmt"
       "log"

       "github.com/kelseyhightower/envconfig"
    )

### Package main

    package main

## Variables

	var name string

	var (
		name string
		age  int
		cool bool
	)

	func main() {
		name := "Kelsey Hightower"
		var age = 32
	}

## Arrays

    var buffer [256]byte
    len(buffer)

    b := [2]string{"Penn", "Teller"}

	// Let the compiler count the elements
	b := [...]string{"Penn", "Teller"} 


## Slices

    var slice []byte = buffer[100:150]
    slice := buffer[100:150]

	names = make([]string, 0, 100)
	names := []string{
		"Rob",
		"Robert",
		"Ken",
	} 

Interation:

	for i, name := range names {
		fmt.Printf("Name: %s Index: %d", name, i) 
	}

## Maps

	map[string]string
	map[string]interface{}

	type Person struct {
		Name string
		Age int
		Cool bool
	}

	m := make(map[string]Person)
	m["me"] = Person{"Kelsey", 33, true}

	p, ok := m["me"]
	if !ok {
		// I don't exist
	}

## Errors

	import "errors"

	errors.New("cannot read file from disk")


Using the "fmt" package.


	import "fmt"

	fmt.Errorf("got an error: %s", err)


Handling errors.

	err := DoSomething()
	if err != nil {
		// Handle error
	}

	if err := DoSomething(); err != nil {
		// Handle
	}

## Structs

	type Person struct {
		Name string
		Age  int
		Cool bool
	}


### Methods

	func (p *Person) ChangeJobTitle(title string) error {
		p.jobTitle = title
	}

