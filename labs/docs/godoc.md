# Documenting with GoDoc

## Clone the Pinger Repo

#### Change Directory

    cd ${GOPATH}/src

#### Run

    git clone https://github.com/kelseyhightower/pinger.git

## Add Comments

#### Change Directory

    cd ${GOPATH}/src/pinger

#### Edit

    ping/pinger.go

-

	package ping

	import (
		"fmt"
		"net/http"
		"time"
	)

	// Objects implementing the Pinger interface can be used to contact
    // a particular website and return roundtrip results.
	type Pinger interface {
		Ping() (*Result, error)
	}

	// Result represents the results of a website ping.
	type Result struct {
		Url      string
		Duration time.Duration
	}

	// Target represents a target website.
	type Target struct {
		url string
	}

	// NewTarget returns a *Target
	func NewTarget(url string) *Target {
		return &Target{url}
	}

	// Ping makes an HTTP GET request to a target identified by t.url
	// and records total roundtrip duration.
	func (t *Target) Ping() (*Result, error) {
		startTime := time.Now()
		res, err := http.Get(t.url)
		duration := time.Since(startTime)

		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("ping error from %s bad status: %d", t.url, res.StatusCode)
		}
		return &Result{t.url, duration}, nil
	}

#### Edit

    ping/doc.go 

-

    // Package ping provides Pinger.
    package ping

#### Edit

    ping/example_test.go

-

	package ping_test

	import (
		"fmt"

		. "pinger/ping"
	)

	func ExampleNewTarget() {
		t := NewTarget("http://google.com")
		res, err := t.Ping()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s - %s", res.Url, res.Duration)
	}

#### Run

    godoc -http=":6060"

#### Visit

    http://localhost:6060/pkg
