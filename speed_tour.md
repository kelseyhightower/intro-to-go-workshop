# Speed Tour

* Statically compiled
* Fast like C++ and Java
* Less memory than Java, Ruby, and Python
* Large standard library
* Small language like C
* Concurrency build-in — No Unicorn or multiple process required
* Garbage collection
* Great for building web services, worker process, and command line tools
* Single binary — no gems or dynamic libs*


### Variables

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

### Arrays

* type definition specifies a length and an element type
* size is fixed. Length is part of the type.
* an array variable denotes the entire array. When ever you pass it around, you copy it.
* To save on the copy you can pass a pointer to the array.
* like a struct but indexed vs named fields. It's a fixe-sized composite value
* Foundation for other types
* Hardly ever used
* 0 indexed
* Accessing the Array past the index will crash your program
* Most common use case is to hold storage for a slice

    // represents an array of four integers
    [4]int 

    var buffer [256]byte
	buffer[0]
    buffer[1]
    len(buffer)  // 256

    // Creating an array.
	b := [2]string{"Penn", "Teller"}

	// Let the compiler count the elements
	b := [...]string{"Penn", "Teller"} 

	// The type of b is [2]string


### Slices

* A slice is backed by an array
* No specified length
* declared just like an array literal, execpt you leave out the element count
* sequence of typed data
* abstration built on top of Go's array
* Grow a slice using the built-in append() function
* Uninitialized slice == nil
* Initialize using built-in make()
* type, length, and capacity
* len(), cap()
* default value is 0. len() and cap() return 0 for a nil slice

    var slice []byte = buffer[100:150]
    slice := buffer[100:150]

The slice data structure.

* slice header

    []string
    []byte

    var names []string

	names = make([]string, length, capacity)
	names := []string{
		"Rob",
		"Robert",
		"Ken",
	} 

	for i, name := range names {
		fmt.Printf("Name: %s Index: %d", name, i) 
	}

### Structs


	type Person struct {
		Name string
		Age  int
		Cool bool
	}


### Methods

	func (p *Person) ChangeJobTitle(title string) error {
		p.jobTitle = title
	}


### Maps

	map[string]string
	map[string]interface{}

	type Person struct {
		Name string
		Age int
		Cool bool
	}

	make(map[string]Person)


### Errors

	import "errors"

	errors.New("cannot read file from disk")


Using the "fmt" package.


	import "fmt"

	fmt.Errorf("got an error: %s", err)


Handling errors.

	err := DoSomething()
	if err != nil {
		// Handle error
		return err
	}
