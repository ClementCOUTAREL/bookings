package main

import (
	"net/http"

	"github.com/coutarel/bookings/internal/config"
	"github.com/coutarel/bookings/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/rooms/general", handlers.Repo.Generals)
	mux.Get("/rooms/major", handlers.Repo.Majors)
	mux.Get("/reservation-availability", handlers.Repo.ReservationAvailability)
	mux.Post("/reservation-availability", handlers.Repo.PostReservationAvailability)
	mux.Get("/make_reservation", handlers.Repo.MakeReservation)
	mux.Post("/make_reservation", handlers.Repo.PostMakeReservation)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
