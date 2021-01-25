package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Mux is a request handler -  It matches the URL of each incoming request against a list of registered patterns
	//and calls the handler for the pattern that most closely matches the URL.

	//servemux supports two URL patterns
	//fixed - dont end with / - /snippet
	//subtree - end with / - like root - catches everything underneath it
	//HandleFunc - register routes with DefaultServeMux (regular servermux) behind the scenes - var DefaultServeMux = NewServeMux()
	//host patterns processed first - foo.example.org/ before /baz
	//longer url patterns get priority over short url patterns

	//DefaultServeMux is a global var so other 3rd party packages can access it  if compromised = security risk
	//thus better to create own locally-scoped servemux instead
	mux := http.NewServeMux()
	//route requests to root to home handler
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("Starting server on :4000")
	//address should be in format :port
	//only specific host if machine has multi network interfaces
	//if you use :name (:http) go will look up port number from /ect/services
	err := http.ListenAndServe(":4000", mux)

	//log.fatal calls os.Exit(1) after writing message thus causing app to exit immediately
	log.Fatal(err)
}
