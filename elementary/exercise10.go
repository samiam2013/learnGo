package main

import "fmt"

// Write a program that prints the next 20 leap years
func main() {
	leapYear := 2024
	for i := 0; i < 20; i++ {
		fmt.Println(leapYear, "is a leap year.")
		leapYear += 4
	}
}
