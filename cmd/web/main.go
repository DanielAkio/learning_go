package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DanielAkio/learning_go/internal/config"
	"github.com/DanielAkio/learning_go/internal/handlers"
	"github.com/DanielAkio/learning_go/internal/models"
	"github.com/DanielAkio/learning_go/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = "8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on localhost:%s", portNumber)
	srv := &http.Server{
		Addr:    ":" + portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// command to generate go.mod file
// go mod init github.com/DanielAkio/go-course

// commands
// go run ./cmd/web
// go get -u github.com/go-chi/chi/v5
// go mod tidy
