# pinger

Log response times from a list of websites

#### Create

	${GOPATH}/src/pinger

#### Edit

	${GOPATH}/src/pinger/main.go

-

	package main

	import (
		"fmt"
		"log"
		"net/http"
		"sync"
		"time"
	)

	type Pinger interface {
		Ping() (*Result, error)
	}

	type Result struct {
		Url      string
		Duration time.Duration
	}

	type Target struct {
		url string
	}

	func NewTarget(url string) *Target {
		return &Target{url}
	}

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

	func pinger(work chan Pinger, result chan *Result, wg *sync.WaitGroup) {
		for {
			w, ok := <-work
			if !ok {
				break
			}
			res, err := w.Ping()
			if err != nil {
				log.Printf("pinger: %s", err)
				continue
			}
			result <- res
		}
		close(result)
		wg.Done()
	}

	func printer(result chan *Result, wg *sync.WaitGroup) {
		for {
			res, ok := <-result
			if !ok {
				break
			}
			log.Printf("ping %s %s", res.Url, res.Duration)
		}
		wg.Done()
	}

	func main() {
		targets := []string{
			"http://google.com",
			"http://puppetlabs.com",
			"http://newrelic.com",
		}
		work := make(chan Pinger)
		result := make(chan *Result)

		var wg sync.WaitGroup
		wg.Add(2)
		go pinger(work, result, &wg)
		go printer(result, &wg)

		for _, t := range targets {
			work <- NewTarget(t)
		}
		close(work)
		wg.Wait()
	}

