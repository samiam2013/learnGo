package elementary

import "fmt"

// Ex02 exports GreetName redundantly. sorry.
func Ex02() {
	GreetName()
}

// GreetName : 2. Write a program that asks the user for their name and greets them with their name.
func GreetName() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)
	fmt.Print("\nNice to meet you " + name)
}
