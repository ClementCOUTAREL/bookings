package models

import "github.com/coutarel/bookings/internal/forms"

// Template data holds data sent from handlers to the templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Error     string
	Warning   string
	Form      *forms.Form
}
