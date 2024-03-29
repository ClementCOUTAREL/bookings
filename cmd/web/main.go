package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/coutarel/bookings/internal/config"
	"github.com/coutarel/bookings/internal/handlers"
	"github.com/coutarel/bookings/internal/models"
)

const portNumber = "localhost:8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	cache, err := handlers.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create a template cache")
	}

	app.TemplateCache = cache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	handlers.NewTemplates(&app)

	return nil
}
