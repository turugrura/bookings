package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/turugrura/bookings/internal/config"
	"github.com/turugrura/bookings/internal/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

var functions = template.FuncMap{}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)

	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache.")
	}

	// buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	// _, err := buf.WriteTo(w)
	// if err != nil {
	// 	log.Println("Error writing to browser.", err)
	// }

	err := t.Execute(w, td)
	if err != nil {
		log.Println("Error writing to browser.", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	templatePath := "./templates"

	pages, err := filepath.Glob(templatePath + "/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		maches, err := filepath.Glob(templatePath + "/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(maches) > 0 {
			ts, err = ts.ParseGlob(templatePath + "/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
