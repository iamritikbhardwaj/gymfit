package main

import (
	"net/http"
)

// the server error helper writes a log entry at the error level (Including the request methods and uri as attributesA)
// then sends a generic 500 error to the user

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

 // the clientError helps in sending client specefic error code and corresponding description 
 func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
 }