package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/coutarel/bookings/pkg/config"
	"github.com/coutarel/bookings/pkg/models"
)

// Repo is the repository variable for the handlers package
var Repo *Repository

// Repository is the repository definition struct
type Repository struct {
	App *config.AppConfig
}

// NewRepo initializes the repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers package
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the about page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "ip", remoteIp)
	renderTemplate(w, r, "home.page.go.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello test map"
	stringMap["ip"] = repo.App.Session.GetString(r.Context(), "ip")

	renderTemplate(w, r, "about.page.go.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (repo *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "generals.page.go.tmpl", &models.TemplateData{})
}

func (repo *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "majors.page.go.tmpl", &models.TemplateData{})
}

func (repo *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "make_reservation.page.go.tmpl", &models.TemplateData{})
}

func (repo *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "contact.page.go.tmpl", &models.TemplateData{})
}

func (repo *Repository) ReservationAvailability(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "reservation.page.go.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (repo *Repository) PostReservationAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	if start > end {
		log.Println("Wrong dates")
	}

	res := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(out)
}
