# Managing Dependencies with Godep

## Install Godep

#### Change Directory

    cd ${HOME}

#### Run

    go get github.com/tools/godep

## Save Dependencies

#### Change Directory

    cd ${GOPATH}/src/github.com/${username}/csv2json-cli

#### Inspect

    go list -json

#### Vendor

    godep save

#### Inspect 

    cat Godeps/Godeps.json

#### Inspect

    ls -R Godeps

#### Version

    git add Godeps
    git commit -m "Manage dependencies with Godep"


## Building with Godep

#### Run

    godep go build -o csv2json .

## Testing with Godep

#### Run

    godep go test
