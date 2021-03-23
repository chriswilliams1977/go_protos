package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	//if this is not being picked up make sure all referenced package files have no errors
	//if they have errors GOLand will not pull in mod reference correctly
	snippets "github.com/chriswilliams1977/protos/internal/app/snippets"
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

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Get static files path and set to defaul if no value is passed
	uiPath := snippets.GetEnv("UI_PATH","./web/")

	// Initialize a new instance of application containing the dependencies.
	app := &snippets.Application{
		ErrorLog: errorLog,
		InfoLog: infoLog,
		UiPath: uiPath,
	}

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.Routes(), // Call the new app.routes() method
	}

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	//address should be in format :port
	//only specific host if machine has multi network interfaces
	//if you use :name (:http) go will look up port number from /ect/services
	infoLog.Printf("Starting server on %s", *addr)

	//The http.ListenAndServe() function takes a http.Handler object as the second
	//func ListenAndServe(addr string, handler Handler) error
	//parameter but mux satisfies the http.Handler interface thus can pass mux
	//servemux is just being a special kind of handler,
	//which instead of providing a response itself passes the request on to a second handler
	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it.
	//err := http.ListenAndServe(*addr, mux)
	//this uses the srv struct so we can append custom error logging
	//By default, if Go’s HTTP server encounters an error it will log it using the standard logger
	err := srv.ListenAndServe()

	//log.fatal calls os.Exit(1) after writing message thus causing app to exit immediately
	//log.Fatal(err)
	errorLog.Fatal(err)
}