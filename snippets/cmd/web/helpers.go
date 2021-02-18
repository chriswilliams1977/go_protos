package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"os"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {

	//debug.Stack() is used to get current goroutine stack trace and append to log message
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	//set call depth to get filename and line number
	app.errorLog.Output(2, trace)

	//http.StatusInternalServerError is a constant for 500 instead of writing it out
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {

	//http.StatusText creates a human friendly version of HTTP status code
	http.Error(w, http.StatusText(status), status)

}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func (app *application) notFound(w http.ResponseWriter) {

	app.clientError(w, http.StatusNotFound)

}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
