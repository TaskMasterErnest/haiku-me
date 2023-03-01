package main

import (
	"net/http"

	"github.com/justinas/alice"
)

//the routes method returns a servemux containing our application routes
func (app *application) routes() http.Handler {
	// prepping a server
	mux := http.NewServeMux()

	//serving the static files for frontend
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	//handling all server functions
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}
