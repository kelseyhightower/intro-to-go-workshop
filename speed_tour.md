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
		name = "Kelsey Hightower"
		distro := "CoreOS"
		fmt.Printf("Name: %s\nLocation: %s\nDistro: %s\n", name, Location, distro)
	}


## Arrays

    func main() {
		locations := [3]string{
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

## Functions

    func ping(url string) bool {
		resp, err := http.Get(string)
		if err != nil {
			return false
		}
		if resp.StatusCode != http.StatusOK {
			return false
		}
		return true
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
