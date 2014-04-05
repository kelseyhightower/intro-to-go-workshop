# Speed Tour

## Packages

    import(
       "fmt"
       "log"

       "github.com/kelseyhightower/targz"
    )

## Variables

	var (
		name     string
    	Location = "Portland"
	)

	func main() {
    	var name string = "Kelsey Hightower"

        var name string
		name = "Kelsey Hightower"

        name := "Kelsey Hightower"
	}

## Arrays

    func main() {
    	var buffer [256]byte

		locations := [...]string{
    		"Long Beach",
    		"Atlanta",
    		"Portland",
    	} 

    	fmt.Printf("Number of locations: %d", len(locations))
    }

## Slices

	func main() {
		locations := []string{
			"Long Beach",
			"Atlanta",
			"Portland",
			"New York",
			"Denver",
			"Dallas",
		}
		for _, name := range locations {
			fmt.Printf("Name: %s\n", name)
		}

		middleTwo := locations[2:4]
		fmt.Printf("%#v", middleTwo)

        ints := make([]int, 0)
    	for i := 0; i <= 100; i++ {
    		ints = append(ints, i)
		}
    	fmt.Printf("%#v", ints)
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
		name     string
		Location string
	}

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

## Channels and Goroutines

	func doubler(input, output chan int) {
		for {
			i := <-input
			output <- i * 2
		}
	}

	func printer(output chan int) {
		for {
			fmt.Printf("%d\n", <-output)
		}
	}

	func main() {
		input := make(chan int)
		output := make(chan int)

		go doubler(input, output)
		go printer(output)

		for i := 0; i <= 10; i++ {
			input <- i
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
