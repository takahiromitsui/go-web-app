package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/takahiromitsui/go-web-app/pkg/config"
	"github.com/takahiromitsui/go-web-app/pkg/handlers"
	"github.com/takahiromitsui/go-web-app/pkg/render"
)

const port = ":8000"

func main() {
	var app config.AppConfig
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // session cookie persists after the browser is closed => later store it in a database
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false // set to true in production

	tc, err := render.CreateTemplateToCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplatesCache = tc
	app.UseCache = false
	render.NewTemplate(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println("Server running on port", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

