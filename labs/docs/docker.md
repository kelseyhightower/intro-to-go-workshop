# Deploy with Docker

## Install Docker

Installation Instructions: http://docs.docker.io/en/latest/installation

#### Change Directory

    cd ${HOME}

#### Download

    curl -o docker https://get.docker.io/builds/Linux/x86_64/docker-latest

#### Install
	
	chmod +x docker
    sudo cp docker /usr/local/bin/

#### Configure

    export DOCKER_HOST=tcp://${docker_host}:4243


## Create a Dockerfile

#### Change Directory

    cd ${GOPATH}/src/github.com/${username}/csv2json-server

#### Edit

    Dockerfile

-

	# csv2json-server
	#
	# VERSION 0.0.1

	FROM       ubuntu
	EXPOSE     8080
	ADD        csv2json-server /usr/local/bin/csv2json-server
	ENTRYPOINT ["/usr/local/bin/csv2json-server"]

#### Build

    GOOS=linux go build -o csv2json-server .

## Build the Docker Image

#### Run

    docker build -t ${username}/csv2json-server .

#### Inspect

    docker images

## Run the Docker Image

#### Run

    docker run -d -P ${username}/csv2json-server

## Testing with curl

#### Inspect

    docker port ${container} 8080

#### Run

    curl -X POST http://${docker_host}:${port}/csv2json --data-binary @${HOME}/famous-gophers.csv

