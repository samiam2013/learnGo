package main

import "fmt"

func exercise02() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)
	fmt.Print("\nNice to meet you " + name)
}
