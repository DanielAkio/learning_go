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
}

func lessonOne() {
	var isTrue bool

	isTrue = false

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
}

func lessonTwo() {
	cat := "cat"

	if cat == "cat" {
		log.Println("It's a cat")
	} else {
		log.Println("It's not a cat")
	}
}

func lessonThree() {
	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is false")
	} else if myNum < 100 && isTrue {
		log.Println("myNum is less than 100 and isTrue is true")
	} else if myNum == 101 || isTrue {
		log.Println("myNum is 101 or isTrue is true")
	} else if myNum > 1000 && isTrue == false {
		log.Println("myNum is greater than 1000 and isTrue is false")
	}
}

func lessonFour() {
	myVar := "cat"

	// It only run the case that matches
	switch myVar {
	case "dog":
		log.Println("It's a dog")
	case "cat":
		log.Println("It's a cat")
	case "fish":
		log.Println("It's a fish")
	default:
		log.Println("It's not a dog, cat or fish")
	}
}
