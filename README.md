# Bird Encyclopedia - A Full Stack Web Application built using Golang

## Aim:

To build a community encyclopedia of birds.
This website will :

- Display the different entries submitted by the community, with the name and details of the bird they found.
- Allow anyone to post a new entry about a bird that they saw.

## Components:

- Web Server
- Front-end(Client side) app
- Database

## Implementation steps:
### Part1: Creating a basic server

- mkdir birdpedia
- go mod init ./birdpedia
- touch main.go
- Write main function
- go run main.go
- curl localhost:8080 or go to http://localhost:8080
- try 8080/(random string) - will give same result

### Part2:

- Make changes using gorilla/mux
- go run main.go
- curl localhost:8080 or go to http://localhost:8080 - will give 404.
- Try 8080/hello - works..everything else gives 404

Write test file to test the functions

- To run this test, just run :
  go test ./...

Initially,the test was working for any URL.
By using our router, we are testing for our specific URL: 8080/hello
I wrote tests for both GET and POST methods.

- To perform the tests use:
  go test ./...

### Part 3: Adding static files

- mkdir assets
- touch assets/index.html
- Modify router to read static files
- Add tests for the static file server
- To test your server use: go test ./...
- If you want your succes messages to be printed as welll, use the -v verbose flag:
  go test -v ./...

### Part 4: Adding REST APIs

- Create tables and forms in the html file.
  The form uses POST to that will add an entry to our existing list of birds
- The table is populated by fetching the list of all birds currently in the system by using GET
- Now write handlers for the post and get requests in the bird_handlers.go file.
- Update the router in main.go to handle these requests.
- Run the program using: go run main.go bird_handlers.go
- If your handlers are in the main function as well, it is sufficient if you only run main.go. As I added the bird_handlers in a different file for code modularity, we need to run both the files.
- At this point our application will be able to reflect the latest addition to the birds table but it won't be able to store the data.

Add test files for testing the current existing functionality and run the tests using go run main.go bird_handlers.go

### Part 5: Adding PostgreSQL database

- Open postgres terminal
- Create database using command: CREATE DATABASE bird_encyclopedia;
- Use command: \l to check if the database is created successfully
- Connect to it using command: \c bird_encyclopedia
- CREATE TABLE birds (
  id SERIAL PRIMARY KEY,
  species VARCHAR(256),
  description VARCHAR(1024)
  );

- Import package for running postgres commands using  go get "github.com/lib/pq"..Add this package to imports in files making postgres connection.
- Create the store.go file in our current directory to interact with the database
- Update our bird handlers to send and receive data to our database
- Use go get "github.com/stretchr/testify/suite" to use in creating the test suite
- Add test files for testing the database store
Make sure that the database names and table names match correctly.
- Perform tests to check if everything works correctly.
- Now run the program using:
 go run main.go bird_handlers.go store.go   
- Add any new birds and they should be visible in table even after reloading.


### Part 6: Making the Postgres database on a docker container
- Open docker
- Run command: docker run --name postgres-container -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres
- Create Dockerfile anddatabase.sql file. 
- Save the above two files in a directory within the user directory
- cd to the new directory and run docker build -t my-postgres-img ./
- Check if the image is properly created by using the command: docker images -a
- Now run the image using the command:docker run -d --name my-postgresdb-container -p 5432:5432 my-postgres-db
- This image contains our database-bird_encyclopedia which contains our birds table because of the sql file dump.




