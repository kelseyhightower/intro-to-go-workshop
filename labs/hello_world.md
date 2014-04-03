# Hello World

### Create

```
mkdir ${GOPATH}/src/github.com/${username}/hello
```

### Edit

`${GOPATH}/src/github.com/${username}/hello/main.go`

### Code

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
```

### Build

```
go build main.go -o hello
```

### Run

```
./hello
```
