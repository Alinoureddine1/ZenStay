package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Alinoureddine1/ZenStay/internal/config"
	"github.com/Alinoureddine1/ZenStay/internal/forms"
	"github.com/Alinoureddine1/ZenStay/internal/models"
	"github.com/Alinoureddine1/ZenStay/internal/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About renders the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Test string"

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// RoyalSuites renders the royal suites page
func (m *Repository) RoyalSuites(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "royal-suite.page.tmpl", &models.TemplateData{})
}

// DeluxeSuites renders the deluxe suites page
func (m *Repository) DeluxeSuites(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "deluxe-suite.page.tmpl", &models.TemplateData{})
}

// SearchAvailability renders the search availability page
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte("start date is " + start + " and end date is " + end))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		m.App.ErrorLog.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

// Reservation renders the make a reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{

		Form: forms.New(nil),
	})
}

// PostReservation handles posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
