package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

// var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	birds, err1 := store.GetBirds()

	birdListBytes, err := json.Marshal(birds) //convert birds variable to json
	if err1 != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := Bird{}
	// We send all our data as HTML form data the `ParseForm` method of the request, parses the form values
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")
	// birds = append(birds, bird)
	err = store.CreateBird(&bird)
	if err != nil {
		fmt.Println(err)
	}

	// redirect the user to the original HTMl page located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
