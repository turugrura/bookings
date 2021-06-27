package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/turugrura/bookings/pkg/config"
	"github.com/turugrura/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

var functions = template.FuncMap{}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// data for every page

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
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

	td = AddDefaultData(td)

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
