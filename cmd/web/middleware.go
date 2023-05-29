package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if user is not authenticated, redirect them to login page and return without executing any more handlers in chain
		if !app.isAuthenticated(r) {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}

		//otherwise, set the cache so that pages that require authentication are not stored in browser cache
		w.Header().Add("Cache-Control", "no-store")

		//call the next handler in chain
		next.ServeHTTP(w, r)
	})
}

//a NoSurf middleware function to guard against CSRF
func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//retrieve authentication data from the session data, if not available, call the next handler in the chain
		id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
		if id == 0 {
			next.ServeHTTP(w, r)
			return
		}

		//otherwise check to see if a user with that ID exists
		exists, err := app.users.Exists(id)
		if err != nil {
			app.serverError(w, err)
			return
		}

		//if a matching user is found, a copy of the request context is created.
		if exists {
			ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
			r = r.WithContext(ctx)
		}

		//call the next handler in the chain
		next.ServeHTTP(w, r)

	})
}
