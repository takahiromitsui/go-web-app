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
	render.RenderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}