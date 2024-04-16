package main

import (
	"html/template"
)

func renderTemplate(file string) *template.Template {
	templ := template.Must(template.ParseFiles("./src/public/" + file))
	return templ
}
