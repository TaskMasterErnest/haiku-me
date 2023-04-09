package main

import (
	"net/http"

	"github.com/TaskMasterErnest/ui"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

//the routes method returns a servemux containing our application routes
func (app *application) routes() http.Handler {
	// prepping a server
	router := httprouter.New()

	//serving the static files for frontend
	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	//creating a new dynamic middleware to cater to the new sessionManager middleware application
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	//update the routes to use the dynamic middleware so as to farm session info
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	//new routes to be used by the dynamic middleware
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	// the protected routes, with the new protected middleware chain
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
