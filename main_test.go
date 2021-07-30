package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	r := newRouter()                    //// Instantiate the router
	mockServer := httptest.NewServer(r) //The mock server we created runs a server and exposes its location in the URL attribute

	//Test 1: for get method and actual url
	resp, err := http.Get(mockServer.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK { //we need 200
		t.Errorf("status should be ok, we got %d", resp.StatusCode)
	}

	//To read response body and convert to a string
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	respString := string(bytes)
	expected := "Hello World! Welcome to Sudeeptha's Bird Encyclopedia"
	if respString != expected {
		t.Errorf("rrsponse should be %s we got %s", expected, respString)
	}

	//Test 2: for post method
	resp, err = http.Post(mockServer.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We want our status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// The code to test the body is also mostly the same, except this time, we expect an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString = string(b)
	expected = ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}
