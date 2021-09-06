package elementary

import "fmt"

func Exe02() {
	GreetName()
}

func GreetName() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)
	fmt.Print("\nNice to meet you " + name)
}
