package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Implementation part1-//in this case whatever route we give, it will give standard response
// The "net/http" library has methods to implement HTTP clients and servers
// Both "fmt" and "net" are part of the Go standard library

// func main() {
// 	http.HandleFunc("/", handler) // The "HandleFunc" method accepts a path and a function as arguments. But handler should have have a proper signature
// 	// After defining our server, we finally "listen and serve" on port 8080
// 	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
// 	// and the handler defined above (in "HandleFunc") is used
// 	http.ListenAndServe(":8080", nil)

// }

// // "handler" is our handler function. It has (ResponseWriter, Request type) as the arguments.
// func handler(writer http.ResponseWriter, req *http.Request) {

// 	// fmt.Fprintf takes a “writer” as its first argument. The second argument is the data that is piped into this writer.
// 	// The output therefore appears according to where the writer moves it.
// 	// In our case the ResponseWriter w writes the output as the response to the users request.
// 	fmt.Fprintf(writer, "Hello. Welcome to Sudeeptha's Bird Encyclopedia")
// }

//Implementation part2 - in this case the router will work only for 8080/hello..everything else will return 404

func newRouter() *mux.Router {
	r := mux.NewRouter()                           //creating new router
	r.HandleFunc("/hello", handler).Methods("GET") //router helps us specify which method the path will be va;id for
	return r
}
func main() {
	r := newRouter()                //creating new router
	http.ListenAndServe(":8080", r) //passing our router after declaring our routes
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! Welcome to Sudeeptha's Bird Encyclopedia")
}
