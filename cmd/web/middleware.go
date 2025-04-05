package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurve adds CSRF protection to all POST requests
func NoSurve(next http.Handler) http.Handler {
	csrHandler := nosurf.New(next)

	csrHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProducion,
		SameSite: http.SameSiteLaxMode,
	})
	return csrHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
