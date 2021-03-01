package main

import (
	"fmt"

	"Github.com/Samiam2013/GoExercise/moduleExample"
)

func main() {
	input := 15
	fmt.Printf("input for fizzbuzz: %v \n", input)
	output := moduleExample.FizzBuzz(15)
	fmt.Println(output)
}
