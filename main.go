package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8000"

type Film struct {
	Title string
	Director string
}

func renderTemplate(file string) *template.Template {
	templ := template.Must(template.ParseFiles("templates/" + file))
	return templ
}


func Home(w http.ResponseWriter, r *http.Request) {
	films := map[string][]Film{
		"films": {
			{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "The Dark Knight", Director: "Christopher Nolan"},
		},
	}
	renderTemplate("home.page.html").Execute(w, films)
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate("about.page.html").Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, nil)
}

