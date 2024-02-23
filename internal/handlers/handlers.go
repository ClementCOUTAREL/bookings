package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/coutarel/bookings/internal/config"
	"github.com/coutarel/bookings/internal/forms"
	"github.com/coutarel/bookings/internal/models"
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
	repo.App.Session.Get(r.Context(), "error")
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	renderTemplate(w, r, "make_reservation.page.go.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
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

// PostReservation handles the posting of a reservatrion form
func (repo *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.MinLength("last_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		renderTemplate(w, r, "make_reservation.page.go.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}

	repo.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (repo *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := repo.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("Cannot get item from session")
		repo.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	repo.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	renderTemplate(w, r, "reservation_summary.page.go.tmpl", &models.TemplateData{
		Data: data,
	})
}
