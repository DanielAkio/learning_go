package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name          string `json:"Name"`
	Age           int    `json:"Age"`
	Job           string `json:"Job"`
	LicenseToKill bool   `json:"LicenseToKill"`
}

func main() {
	log.Println("--- lessonOne")
	lessonOne()
	log.Println("--- lessonTwo")
	lessonTwo()
}

func lessonOne() {
	myJson := `
	[
		{
			"Name": "James Bond",
			"Age": 32,
			"Job": "Secret Agent",
			"LicenseToKill": true
		},
		{
			"Name": "Miss Moneypenny",
			"Age": 27,
			"Job": "Secretary",
			"LicenseToKill": false
		}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)
}

// write json from a struct
func lessonTwo() {
	var mySlice []Person

	var m1 Person
	m1.Name = "Wallace"
	m1.Age = 30
	m1.Job = "Teacher"
	m1.LicenseToKill = false
	mySlice = append(mySlice, m1)

	var m2 Person
	m2.Name = "Gromit"
	m2.Age = 22
	m2.Job = "Student"
	m2.LicenseToKill = false
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
	}

	fmt.Println(string(newJson))
}
