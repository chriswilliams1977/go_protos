package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {

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
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	//./ui/static/
	fileServer := http.FileServer(http.Dir(app.uiPath+"static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}