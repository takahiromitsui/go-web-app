package main

import (
	"html/template"
)

func renderTemplate(file string) *template.Template {
	templ := template.Must(template.ParseFiles("templates/" + file))
	return templ
}
