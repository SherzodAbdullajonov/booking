package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

//NoSurf adds CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrFHandler := nosurf.New(next)
	csrFHandler.SetBaseCookie(http.Cookie{
		HttpOnly: false,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrFHandler
}

//SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return Session.LoadAndSave(next)
}
