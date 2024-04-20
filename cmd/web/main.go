package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/takahiromitsui/go-web-app/pkg/config"
	"github.com/takahiromitsui/go-web-app/pkg/handlers"
	"github.com/takahiromitsui/go-web-app/pkg/render"
)

const port = ":8000"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateToCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplatesCache = tc
	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, nil)
}

