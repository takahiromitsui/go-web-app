package handlers

import (
	"net/http"

	"github.com/takahiromitsui/go-web-app/pkg/render"
)


type Film struct {
	Title string
	Director string
}

func Home(w http.ResponseWriter, r *http.Request) {
	films := map[string][]Film{
		"films": {
			{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "The Dark Knight", Director: "Christopher Nolan"},
		},
	}
	render.RenderTemplate("home.page.html").Execute(w, films)
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate("about.page.html").Execute(w, nil)
}