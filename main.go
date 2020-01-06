package main

import (
	"fmt"
	"net/http"
)


// This function executes anytime someone visits our web app
// ResponseWriter is part of http package and an interface we take in
// *http.Request is a pointer
// handlerFunc can be named anything actually
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// first argument is where you want it to print to
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc) 
	http.ListenAndServe(":3000", nil) // localhost is implied; nil means use the built-in serve-mux that comes with the built-in http package
	// ListenAndServe starts up the web server
	// unencrypted web traffic listens on Port 80 by default
	
}