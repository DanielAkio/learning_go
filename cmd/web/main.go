package main

import (
	"fmt"
	"net/http"

	"github.com/DanielAkio/go-course/pkg/handlers"
)

const portNumber = "8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on localhost:%s", portNumber)
	_ = http.ListenAndServe(":"+portNumber, nil)
}

// command to generate go.mod file
// go mod init github.com/DanielAkio/go-course

// command to run
// go run ./cmd/web
