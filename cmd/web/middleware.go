package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// SessionLoad generates a cookie that stores session infos
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// NoSurf generates a crsftoken available on every request via a cookie
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
