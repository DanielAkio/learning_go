package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (d *Gorilla) Says() string {
	return "Ugh! Ugh! Ugh!"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 32,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
