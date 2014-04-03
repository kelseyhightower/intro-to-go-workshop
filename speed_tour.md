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

Represents an array of four integers

    [4]int 

    var buffer [256]byte
	buffer[0]
    buffer[1]
    len(buffer)  // 256

Creating an array.
	
    b := [2]string{"Penn", "Teller"}

	// Let the compiler count the elements
	b := [...]string{"Penn", "Teller"} 


### Slices

Creating a slice:

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

### Structs


	type Person struct {
		Name string
		Age  int
		Cool bool
	}


#### Methods

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
