package main

import (
	"fmt"      // "fmt" has methods for formatted I/O operations (like printing to the console)
	"net/http" // The "net/http" library has methods to implement HTTP clients and servers
	// Both "fmt" and "net" are part of the Go standard library
)

func main() {
	http.HandleFunc("/", handler) // The "HandleFunc" method accepts a path and a function as arguments. But handler should have have a proper signature

	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":8080", nil)

}

// "handler" is our handler function. It has (ResponseWriter, Request type) as the arguments.
func handler(writer http.ResponseWriter, req *http.Request) {

	// fmt.Fprintf takes a “writer” as its first argument. The second argument is the data that is piped into this writer.
	// The output therefore appears according to where the writer moves it.
	// In our case the ResponseWriter w writes the output as the response to the users request.
	fmt.Fprintf(writer, "Hello. Welcome to Sudeeptha's Bird Encyclopedia")
}
