package render

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/takahiromitsui/go-web-app/internal/config"
	"github.com/takahiromitsui/go-web-app/internal/models"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request,
	t string, td *models.TemplateData) error {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		// prod
		tc = app.TemplatesCache
	} else {
		// dev
		tc, _ = CreateTemplateToCache()
	}
	tmpl, ok := tc[t]
	// render the template
	if !ok {
		return errors.New("can't get template from cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	err := tmpl.Execute(buf, td)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
var pathToTemplates = "./templates"


func CreateTemplateToCache() (map[string]*template.Template, error) {
	// create a new template cache
	tc := map[string]*template.Template{}
	// get all page templates
	pages, err := filepath.Glob(pathToTemplates + "/*.page.tmpl")
  if err != nil {
		return tc, err
	}
	// loop through all page templates
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tc, err
		}
		// get the base layout template
		matches, err := filepath.Glob(pathToTemplates + "/*.layout.tmpl")
		if err != nil {
			return tc, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(pathToTemplates + "/*.layout.tmpl")
			if err != nil {
				return tc, err
			}
		}
		tc[name] = ts
	}
	return tc, nil
}
