package handlers

import (
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
	renderTemplate(w, "home.page.go.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello test map"
	stringMap["ip"] = repo.App.Session.GetString(r.Context(), "ip")

	renderTemplate(w, "about.page.go.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
