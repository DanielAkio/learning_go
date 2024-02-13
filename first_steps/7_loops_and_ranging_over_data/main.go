package main

import "log"

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
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
}

func lessonTwo() {
	animals := []string{"cat", "dog", "bird", "cow", "fish"}

	for _, animal := range animals {
		log.Println(animal)
	}
}

func lessonThree() {
	animals := map[string]string{
		"cat":  "meow",
		"dog":  "woof",
		"bird": "tweet",
		"cow":  "moo",
		"fish": "bubble",
	}

	for key, value := range animals {
		log.Println(key, "makes", value)
	}
}

func lessonFour() {
	firstLine := "Printing string"
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}
}

func lessonFive() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Doe", "john@smith.com", 30})
	users = append(users, User{"Mary", "Doe", "mary@smith.com", 20})
	users = append(users, User{"Sally", "Doe", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Doe", "alex@smith.com", 17})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, "is", user.Age, "years old")
	}
}
