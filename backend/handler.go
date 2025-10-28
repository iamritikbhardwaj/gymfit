package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"../frontend/dist/index.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		app.serverError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "index", nil)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		app.serverError(w, r, err)
		return
	}
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from gymfit"))
}

// func (app *application) home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Server","Go")
// 	w.Write([]byte("Hello world"))
// }
	