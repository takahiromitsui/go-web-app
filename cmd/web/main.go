package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/takahiromitsui/go-web-app/internal/config"
	"github.com/takahiromitsui/go-web-app/internal/handlers"
	"github.com/takahiromitsui/go-web-app/internal/models"
	"github.com/takahiromitsui/go-web-app/internal/render"
)

const port = ":8000"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

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

func run() error {
		// what am I going to put in the session
		gob.Register(models.Reservation{})
		// change this to true when in production
		app.InProduction = false

		infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		app.InfoLog = infoLog

		errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		app.ErrorLog = errorLog

		session = scs.New()
		session.Lifetime = 24 * time.Hour
		session.Cookie.Persist = true // session cookie persists after the browser is closed => later store it in a database
		session.Cookie.SameSite = http.SameSiteLaxMode
		session.Cookie.Secure = app.InProduction // set to true in production

		app.Session = session

		tc, err := render.CreateTemplateToCache()
		if err != nil {
			log.Fatal("cannot create template cache")
			return err
		}
		app.TemplatesCache = tc
		app.UseCache = false
		repo := handlers.NewRepo(&app)
		handlers.NewHandlers(repo)
		render.NewTemplate(&app)
	return nil
}