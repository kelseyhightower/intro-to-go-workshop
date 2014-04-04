# Deploy with Docker

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

    docker -H tcp://162.222.181.149:4243 build -t ${username}/csv2json-server .
    docker -H tcp://162.222.181.149:4243 images
    docker -H tcp://162.222.181.149:4243 ps
    docker -H tcp://162.222.181.149:4243 run -d -p 8080:8080 ${username}/csv2json-server
 

## Testing with curl

    curl -X POST http://162.222.181.149:8080/csv2json --data-binary @${HOME}/famous-gophers.csv
