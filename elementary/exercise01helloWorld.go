package elementary

import "fmt"

// Ex01 exports HelloWorld redundantly. Sorry.
func Ex01() {
	HelloWorld()
}

// HelloWorld : 1. Write a program that prints ‘Hello World’ to the screen.
func HelloWorld() {
	fmt.Println("Hello World!")
}
