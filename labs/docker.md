# Deploy with Docker

## Install Docker

Installation Instructions: http://docs.docker.io/en/latest/installation

#### Download

    curl -o docker https://get.docker.io/builds/Darwin/x86_64/docker-latest

#### Install
	
	chmod +x docker
    sudo cp docker /usr/local/bin/

#### Configure

    export DOCKER_HOST=tcp://${docker_host}:4243


## Create a Dockerfile

#### PWD

    ${GOPATH}/src/github.com/${username}/csv2json-server

#### Edit

    ${GOPATH}/src/github.com/${username}/csv2json-server/Dockerfile

-

	# csv2json-server
	#
	# VERSION 0.0.1

	FROM       ubuntu
	EXPOSE     8080
	ADD        csv2json-server /usr/local/bin/csv2json-server
	ENTRYPOINT ["/usr/local/bin/csv2json-server"]

#### Run

    GOOS=linux go build -o csv2json-server .
    docker build -t ${username}/csv2json-server .
    docker images
    docker ps
    docker run -d -P ${username}/csv2json-server

## Testing with curl

    curl -X POST http://${docker_host}:8080/csv2json --data-binary @${HOME}/famous-gophers.csv

