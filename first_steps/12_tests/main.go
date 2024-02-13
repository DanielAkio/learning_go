package main

import (
	"errors"
	"log"
)

// go test -v
// go test -cover
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out

func main() {
	result, err := divide(100, 10)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	log.Println("Result: ", result)
}

func divide(x float32, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}

	result = x / y
	return result, nil
}
