package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// func RenderTemplate(file string) *template.Template {
// 	templ := template.Must(template.ParseFiles("./src/public/" + file, "./src/public/base.layout.tmpl"))
// 	return templ
// }

var tc = make(map[string]*template.Template)



func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in the cache
	_, inMap := tc[t]
	if !inMap {
		// need to add the template to the cache
		log.Println("Creating template and adding to cache")
		err = CreateTemplateToCache(t)
		if err != nil {
			log.Println("error creating template cache", err)
		}
	} else {
		log.Println("Using template from cache")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("error executing template", err)
		return
	}
}

func CreateTemplateToCache(t string) error {
	templates := []string{
		fmt.Sprintf("./src/public/%s.page.tmpl", t),
		"./src/public/base.layout.tmpl",
	}
	// parse the template files...
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}