package main

import (
	"log"

	"github.com/danielakio/learning_go/helpers"
)

// go mod init github.com/danielakio/learning_go

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	myVar.TypeNumber = 1
	log.Println(myVar.TypeName, myVar.TypeNumber)
}
