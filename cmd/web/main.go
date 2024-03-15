package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DanielAkio/learning_go/pkg/config"
	"github.com/DanielAkio/learning_go/pkg/handlers"
	"github.com/DanielAkio/learning_go/pkg/render"
)

const portNumber = "8080"

func main() {
	var app config.AppConfig

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
