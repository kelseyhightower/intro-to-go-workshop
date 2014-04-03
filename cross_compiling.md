# Cross Compiling

Build the toolchain

```
GOOS
GOARCH
CGO_ENABLED
```

Run

```
go env
cd /usr/local/go/pkg
ls
linux_amd64
```
	
Run

```
apt-get install gcc
cd /usr/local/go/src
GOOS=linux GOARCH=386 CGO_ENABLED=0 ./make.bash —no-clean
```

```
cd /usr/local/go/pkg
ls
linux_386 linux_amd64 
```

```
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 ./make.bash —no-clean

cd /usr/local/go/pkg
ls
	
darwin_amd64 linux_386 linux_amd64
```

## Cross compile for Window or OS X

```
cd to a project
GOOS=darwin go build .
```
