package elementary

import "fmt"

// GreetBobOrAlice :
func GreetBobOrAlice() {
	var name string
	fmt.Print("What is your name?: ")
	fmt.Scanln(&name)
	if name == "Bob" || name == "Alice" {
		fmt.Println("Hello " + name + "!")
	} else {
		fmt.Println("I don't recognize that name.")
	}
}
