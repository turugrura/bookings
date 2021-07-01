package render

import (
	"net/http"
	"testing"

	"github.com/turugrura/bookings/internal/models"
)

func TestAddDefault(t *testing.T) {
	var td = models.TemplateData{}

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value 123 not found in session.")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var w myWriter
	err = RenderTemplate(&w, r, "home.page.html", &models.TemplateData{})
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(&w, r, "not-existent.page.html", &models.TemplateData{})
	if err == nil {
		t.Error("render template which not exist.")
	}
}

func TestNewTemplate(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/something", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
