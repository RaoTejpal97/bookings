package main

import (
	"net/http"

	"github.com/tejpal/booking/pkg/config"
	"github.com/tejpal/booking/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurve)
	mux.Use(SessionLoad)

	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
