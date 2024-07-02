package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DanielAkio/learning_go/driver"
	"github.com/DanielAkio/learning_go/internal/config"
	"github.com/DanielAkio/learning_go/internal/handlers"
	"github.com/DanielAkio/learning_go/internal/helpers"
	"github.com/DanielAkio/learning_go/internal/models"
	"github.com/DanielAkio/learning_go/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = "8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Printf("Starting application on localhost:%s\n\n", portNumber)
	srv := &http.Server{
		Addr:    ":" + portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.Reservation{})
	gob.Register(models.RoomRestriction{})

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	log.Println("Connection to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=admin password=admin")
	if err != nil {
		log.Fatal("cannot connect to database! Dying...")
	}
	log.Println("Connected to database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}

// command to generate go.mod file
// go mod init github.com/DanielAkio/go-course

// command to install and remove not used packages
// go run ./cmd/web
// go get -u github.com/go-chi/chi
// go mod tidy

// command to run test
// go test -v ./...
// go test -cover ./...
// go test -coverprofile=coverage.out ./...
// go tool cover -html=coverage.out
