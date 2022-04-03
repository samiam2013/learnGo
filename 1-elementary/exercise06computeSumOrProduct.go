package elementary

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// TriangleOrFac : 6 Write a program that asks the user for a number n and gives them the
//  possibility to choose between computing the sum and computing the product of 1,…,n.
func TriangleOrFac() {
	var inputString string
	fmt.Print("enter a number: ")
	fmt.Scanln(&inputString)
	inputNumber, err := strconv.ParseInt(inputString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	var result int64
	for {
		fmt.Print("would you like to compute the (s)um or (p)roduct?: ")
		fmt.Scanln(&inputString)
		if strings.ToLower(inputString) == "s" {
			result = Σ(inputNumber)
			break
		} else if strings.ToLower(inputString) == "p" {
			result = Factorial(inputNumber)
			break
		}
	}
	fmt.Println("Your result: " + strconv.FormatInt(result, 10))
}

// Factorial does the ! math thing
func Factorial(input int64) int64 {
	product := int64(1)
	for input > 0 {
		product = product * input
		//fmt.Println("product " + strconv.FormatInt(product, 10))
		input = input - 1
	}
	return product
}
