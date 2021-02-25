package main

import (
	"fmt"

	"Github.com/Samiam2013/GoExercise/ModulePilgrimage"
)

func main() {
	input := 15
	fmt.Printf("input for fizzbuzz: %v \n", input)
	output := ModulePilgrimage.FizzBuzz(15)
	fmt.Println(output)
}
