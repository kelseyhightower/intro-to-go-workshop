# Hello World

> Replace ${username} with your github username.

```
mkdir ~/go/src/github.com/${username}/hello
```

**Edit** `~/go/src/github.com/${username}/hello/main.go`

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
```

**Build**

```
go build main.go -o hello
```

**Run**

```
./hello
```
