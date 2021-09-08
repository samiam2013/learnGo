package elementary

import (
	"fmt"
)

// Write a program that prints a multiplication table for numbers up to 12
func MultiplicationTable() {
	for row := 1; row <= 12; row++ {
		for column := 1; column <= 12; column++ {
			product := row * column
			fmt.Printf("%4v", product)
		}
		fmt.Print("\n")
	}
}
