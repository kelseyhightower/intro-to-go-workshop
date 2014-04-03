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
