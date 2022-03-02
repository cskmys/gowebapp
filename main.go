package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) { // function run anytime a request arives at the server
	// "http.ResponseWriter" is where you write the response
	// "http.Request" is request object where you have info regarding the request
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>") // the string output is html with heading tag "h1"
}

func main() {
	http.HandleFunc("/", handlerFunc) // "/" is the path after the address
	// when we go to the path, the "servemux" is going to execute the "handlerFunc" function
	http.ListenAndServe(":3000", nil) // ":3000" stands for "localhost:3000" and "nil" as the routing function means use the default routing function
}
