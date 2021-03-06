package snippets

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"log"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
type Application struct {
	ErrorLog *log.Logger
	InfoLog *log.Logger
	UiPath string
}

//handlers are control logic to write HTTP response headers and bodies
//router maps URL request pattern to handler
//webserver listens for incoming requests
//handler
//takes a http.ResponseWriter - provides methods for assembling a HTTP response and sending it to the user
//and a *http.Request - struct which holds information about the current request
func  (app *Application)  home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		//uses notfound helper
		app.notFound(w)
		return
	}

	//w.Write([]byte("Hello from Snippets"))

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		app.UiPath+"html/home.page.tmpl",
		app.UiPath+"html/base.layout.tmpl",
		app.UiPath+"html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	//file path must either be relative to your current working directory, or an absolute path.
	//Notice that we can pass the slice of file paths
	//as a variadic parameter?
	//Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
	}
}

func (app *Application)  showSnippet(w http.ResponseWriter, r *http.Request){
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w,"Display a specific snippet with ID %d",id)
}

func (app *Application)  createSnippet(w http.ResponseWriter, r *http.Request){
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
		//http.Error(w,"Method not allowed",405)
		//return
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
		return
	}

	w.Write([]byte("Create new snippet"))
}