package render

import (
	"html/template"
)

func RenderTemplate(file string) *template.Template {
	templ := template.Must(template.ParseFiles("./src/public/" + file))
	return templ
}
