package main

import (
	"github.com/SherzodAbdullajonov/booking/package/config"
	"github.com/SherzodAbdullajonov/booking/package/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)

	mux.Get("/about", handlers.Repo.About)
	mux.Get("/general-quarter", handlers.Repo.Generals)
	mux.Get("/major-suite", handlers.Repo.Majors)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
