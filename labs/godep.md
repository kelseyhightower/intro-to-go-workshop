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

## Updating Dependencies

#### Run

    godep restore
    go get -u github.com/${username}/csv2json
    godep save
