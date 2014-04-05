# Testing csv2json-server

#### Change Directory

    cd ${GOPATH}/src/github.com/${username}/csv2json-server

#### Edit

    api_test.go

-

	package main

	import (
		"bytes"
		"testing"
		"net/http"
		"net/http/httptest"
	)

	type httpTestResponse struct {
		statusCode int
		body       string
	}

	var csv2JsonServerPostBodyTest = `
	Mac, 1947, The Goofy Gophers
	Tosh, 1947, The Goofy Gophers
	Samuel J. Gopher, 1966, Winnie the Pooh
	Chief Running Board, 1968, Go Go Gophers
	Ruffled Feathers, 1968, Go Go Gophers
	`

	var csv2JsonServerWantBodyTest = `[
	  {
		"Date": "1947",
		"Name": "Mac",
		"Title": "The Goofy Gophers"
	  },
	  {
		"Date": "1947",
		"Name": "Tosh",
		"Title": "The Goofy Gophers"
	  },
	  {
		"Date": "1966",
		"Name": "Samuel J. Gopher",
		"Title": "Winnie the Pooh"
	  },
	  {
		"Date": "1968",
		"Name": "Chief Running Board",
		"Title": "Go Go Gophers"
	  },
	  {
		"Date": "1968",
		"Name": "Ruffled Feathers",
		"Title": "Go Go Gophers"
	  }
	]`

	func TestCsv2JsonServer(t *testing.T) {
		body := bytes.NewBufferString(csv2JsonServerPostBodyTest)
		req, err := http.NewRequest("POST", "http://example.com/csv2json", body)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()
		csv2JsonServer(w, req)

		want :=  httpTestResponse{http.StatusOK, csv2JsonServerWantBodyTest}
		got := httpTestResponse{w.Code, w.Body.String()}
		if got != want {
			t.Errorf("Want: %#v, got %#v", want, got)
		} 
	}

#### Run

    go test -v

Run using godep

    godep go test -v

#### Version

    git add api_test.go
    git commit -m "add api tests"

