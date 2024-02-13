package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	log.Println("--- lessonOne")
	lessonOne()
	log.Println("--- lessonTwo")
	lessonTwo()
	log.Println("--- lessonThree")
	lessonThree()
	log.Println("--- lessonFour")
	lessonFour()
	log.Println("--- lessonFive")
	lessonFive()
}

func lessonOne() {
	myMap := make(map[string]string)

	myMap["dog"] = "Toddy"
	myMap["cat"] = "Tob"

	myMap["dog"] = "Nescau"

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])
}

func lessonTwo() {
	myIntMap := make(map[string]int)

	myIntMap["first"] = 1
	myIntMap["second"] = 2

	log.Println(myIntMap["first"])
	log.Println(myIntMap["second"])
}

func lessonThree() {
	myMap := make(map[string]User)

	myMap["me"] = User{
		FirstName: "Daniel",
		LastName:  "Teixeira",
	}

	log.Println(myMap["me"].FirstName)
}

func lessonFour() {
	mySlice := []int{10, 11, 5, 6}

	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)

	log.Println(mySlice)
	sort.Ints(mySlice)
	log.Println(mySlice)

	log.Println(mySlice[3:6])
}

func lessonFive() {
	names := []string{"Daniel", "Teixeira", "de", "Carvalho"}

	log.Println(names)
}
