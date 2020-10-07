package main

import "fmt"

func main() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)
	fmt.Print("\nNice to meet you " + name)
}
