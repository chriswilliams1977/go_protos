package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

//handlers are control logic to write HTTP response headers and bodies
//router maps URL request pattern to handler
//webserver listens for incoming requests

//handler
//takes a http.ResponseWriter - provides methods for assembling a HTTP response and sending it to the user
//and a *http.Request - struct which holds information about the current request
func home(w http.ResponseWriter, r *http.Request){
	//prevents the catch all for / subtrees
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	w.Write([]byte("Hello from Snippets"))
}

func showSnippet(w http.ResponseWriter, r *http.Request){
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w,"Display a specific snippet with ID %d",id)
}

func createSnippet(w http.ResponseWriter, r *http.Request){
	// Use r.Method to check whether the request is using POST or not.
	// If it's not, use the w.WriteHeader() method to send a 405 status code and
	// the w.Write() method to write a "Method Not Allowed" response body. We
	// then return from the function so that the subsequent code is not executed.
	// Use the Header().Set() method to add an 'Allow: POST' header to the
	// response header map. The first parameter is the header name, and
	// the second parameter is the header value.
	//must call this before w.WriteHeader() or w.Write()
	if r.Method != "POST" {
		w.Header().Set("Allow","POST")
		/*
		//this is an example - in reality you dont call directly you use http.Error
		//you must call w.WriteHeader with response code before w.Write or a 200 will be passed
		//to test use curl -i -X POST http://localhost:4000/snippet/create
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		*/
		//common way to handle response code and pass message to header
		http.Error(w,"Method not allowed",405)
		return
	}

	w.Write([]byte("Create new snippet"))
}

func main(){

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