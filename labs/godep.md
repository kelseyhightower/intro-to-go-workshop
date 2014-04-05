# Managing Dependencies with Godep

## Install Godep

    go get github.com/tools/godep

## Save Dependencies

#### Inspect

    cd ${GOPATH}/src/github.com/${username}/csv2json-cli
    go list -json

#### Vendor

    godep save

#### Inspect 

    cat Godeps/Godeps.json

#### Inspect

    tree Godeps

#### Version

    git add Godeps
    git commit -m "Manage dependencies with Godep"

## Building with Godep

#### Build

    godep go build -o csv2json .

## Testing with Godep

#### Test

    godep go test

## Updating Dependencies

#### Run

    godep restore
    go get -u github.com/${username}/csv2json
    godep save
