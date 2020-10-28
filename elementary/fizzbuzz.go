package main

import "fmt"

// I want to make a contribution to github today,
//  I realized I've never implemented fizzbuzz.
func fizzBuzz() {
	// iterate until overflow?
	for i := 1; i > 0; i++ {
		fmt.Print(i, ": ")
		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}
		fmt.Println()
	}
}
