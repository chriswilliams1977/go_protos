package main

import (
	"fmt"
	"net/http"
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