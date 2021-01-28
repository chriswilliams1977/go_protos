package main

import (
	"flag"
	"log"
	"net/http"
)


func main(){

	//NOTE EACH REQUEST GETS ITS OWN GO ROUTINE = RACE CONDITIONS

	//Mux is a request handler -  It matches the URL of each incoming request against a list of registered patterns
	//and calls the handler for the pattern that most closely matches the URL.

	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	//can also use flag.Int(), flag.Bool() and flag.Float64()
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()

	//servemux supports two URL patterns
	//fixed - dont end with / - /snippet
	//subtree - end with / - like root - catches everything underneath it
	//HandleFunc - register routes with DefaultServeMux (regular servermux) behind the scenes - var DefaultServeMux = NewServeMux()
	//host patterns processed first - foo.example.org/ before /baz
	//longer url patterns get priority over short url patterns
	//DefaultServeMux is a global var so other 3rd party packages can access it  if compromised = security risk
	//thus better to create own locally-scoped servemux instead
	//When our server receives a new HTTP request, it
	//calls the servemux’s ServeHTTP() method. This looks up the relevant handler based on the
	//request URL path, and in turn calls that handler’s ServeHTTP() method.
	mux := http.NewServeMux()

	//route requests to root to home handler
	//The http.HandlerFunc() adapter works by automatically adding a ServeHTTP() method to
	//the home function thus satisfying the http.Handler interface
	//functionally the same as
	//mux := http.NewServeMux()
	//mux.Handle("/", http.HandlerFunc(home))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	//address should be in format :port
	//only specific host if machine has multi network interfaces
	//if you use :name (:http) go will look up port number from /ect/services
	log.Printf("Starting server on %s", *addr)

	//The http.ListenAndServe() function takes a http.Handler object as the second
	//func ListenAndServe(addr string, handler Handler) error
	//parameter but mux satisfies the http.Handler interface thus can pass mux
	//servemux is just being a special kind of handler,
	//which instead of providing a response itself passes the request on to a second handler
	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it.
	err := http.ListenAndServe(*addr, mux)

	//log.fatal calls os.Exit(1) after writing message thus causing app to exit immediately
	log.Fatal(err)
}