# Installing Go

## Installation

#### Change Directory

    cd ${HOME}

### Install a binary distribution
 
#### Run

    wget https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz
    tar -xvf go1.2.linux-amd64.tar.gz -C /usr/local

## Setup the Workspace

### GOPATH

#### Edit

    ${HOME}/.bashrc

-

    export GOPATH="${HOME}/go"
    export PATH="/usr/local/go/bin:${GOPATH}/bin:$PATH"

#### Activate:

    source ${HOME}/.bashrc

### Workspace Directories

#### Create the directories on Unix

    mkdir -p ${GOPATH}/src
    mkdir -p ${GOPATH}/pkg
    mkdir -p ${GOPATH}/bin
    mkdir -p ${GOPATH}/src/github.com/${username}

### Check the Go Environment

    go env

-

	GOARCH="amd64"
	GOBIN=""
	GOCHAR="6"
	GOEXE=""
	GOHOSTARCH="amd64"
	GOHOSTOS="darwin"
	GOOS="darwin"
	GOPATH="/Users/kelseyhightower/go"
	GORACE=""
	GOROOT="/usr/local/go"
	GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
	TERM="dumb"
	CC="clang"
	GOGCCFLAGS="-g -O2 -fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fno-common"
	CXX="clang++"
	CGO_ENABLED="1"
