package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/takahiromitsui/go-web-app/pkg/config"
	"github.com/takahiromitsui/go-web-app/pkg/models"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, t string, td *models.TemplateData) {
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
		log.Fatal("Could not get template from cache")
	}
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}


func CreateTemplateToCache() (map[string]*template.Template, error) {
	// create a new template cache
	tc := map[string]*template.Template{}
	// get all page templates
	pages, err := filepath.Glob("./src/public/*.page.tmpl")
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
		matches, err := filepath.Glob("./src/public/*.layout.tmpl")
		if err != nil {
			return tc, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./src/public/*.layout.tmpl")
			if err != nil {
				return tc, err
			}
		}
		tc[name] = ts
	}
	return tc, nil
}
