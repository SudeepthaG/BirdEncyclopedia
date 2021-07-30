package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()                           //creating new router
	r.HandleFunc("/hello", handler).Methods("GET") //router helps us specify which method the path will be va;id for

	staticFileDirectory := http.Dir("./assets/") //declaring static file directory and pointing it to our static files directory

	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}
func main() {
	r := newRouter()                //creating new router
	http.ListenAndServe(":8080", r) //passing our router after declaring our routes
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! Welcome to Sudeeptha's Bird Encyclopedia")
}
