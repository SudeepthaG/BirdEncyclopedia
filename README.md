# Bird Encyclopedia - A Full Stack Web Application built using Golang

## Aim:

To build a community encyclopedia of birds.
This website will :

- Display the different entries submitted by the community, with the name and details of the bird they found.
- Allow anyone to post a new entry about a bird that they saw.

Components:

- Web Server
- Front-end(Client side) app
- Database

Implementation steps:
mkdir birdpedia
go mod init ./birdpedia
touch main.go
Write main function
go run main.go
curl localhost:8080 or go to http://localhost:8080
try 8080/(random string) - will give same result

Part2:
Make changes using gorilla/mux
go run main.go
curl localhost:8080 or go to http://localhost:8080 - will give 404.
Try 8080/hello - works..everything else gives 404

Write test file to test the functions
To run this test, just run :
go test ./...
Initially,the test ws working for any URL.

By using our router, we are testing for our specific URL: 8080/hello
I wrote tests for both GET and POST methods. To perform the tests use:
go test ./...
