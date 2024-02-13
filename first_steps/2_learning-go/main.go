package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	whatToSay = "Goodbye, cruel world!"
	fmt.Println(whatToSay)

	var i int
	i = 2
	fmt.Println("i is set to", i)

	whatsWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatsWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "Something", "else"
}
