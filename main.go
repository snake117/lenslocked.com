package main // First thing that starts running your program

import (
	"fmt" // used for printing out data to different writers or standard-out
	"net/http" // for handling/creating web requests and running server
)


// This function executes anytime someone visits our web app
// ResponseWriter is part of http package and an interface we take in
// *http.Request is a pointer
// DOES NOT NEED TO BE CALLED handlerFunc, can be named anything actually
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Someone visited our page") prints out in the logs when fresh is running (ie does NOT print in webpage)
	// first argument is where you want it to print to
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc) // we're NOT calling the function, just telling it to use this function
	// as long as URL path matches the PREFIX it will use the handlerFunc


	http.ListenAndServe(":3000", nil) // localhost is implied; nil means use the built-in serve-mux that comes with the built-in http package
	// ListenAndServe starts up the web server and to listen on port 3000
	// If you put localhost:3000, then the request is restricted to only localhost, whereas this way can be used with anything as long it goes through port :3000
	// unencrypted web traffic listens on Port 80 by default
	
}