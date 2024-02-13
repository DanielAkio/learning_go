package main

import "log"

var s = "seven"

func main() {
	log.Println("--- LessonOne")
	lessonOne()
	log.Println("--- LessonTwo")
	lessonTwo()
}

func lessonOne() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}

func lessonTwo() {
	var s2 = "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}
