package handlers

import (
	"fmt"
	"net/http"

	"github.com/turugrura/bookings/pkg/config"
	"github.com/turugrura/bookings/pkg/models"
	"github.com/turugrura/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (a *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (a *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}

func (a *Repository) Notfound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not found.")
}
