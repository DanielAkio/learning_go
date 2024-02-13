package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         string
	BirthDate   time.Time
}

type myStruct struct {
	FirstName string
}

// Receiver
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	log.Println("--- LessonOne")
	lessonOne()
	log.Println("--- LessonTwo")
	lessonTwo()
}

func lessonOne() {
	user := User{
		FirstName:   "Daniel",
		LastName:    "Teixeira",
		PhoneNumber: "+55 11 98211-4798",
	}

	log.Println(user.FirstName, user.LastName)
	log.Println("Birthdate: ", user.BirthDate)
}

func lessonTwo() {
	var myVar myStruct
	myVar.FirstName = "Daniel"

	myVar2 := myStruct{
		FirstName: "Catarina",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
