package main

import (
	"fmt"
	"log"
	"strconv"
)

func exercise05() {
	var inputString string
	fmt.Print("Input a nubmer: ")
	fmt.Scanln(&inputString)
	inputNumber, err := strconv.ParseInt(inputString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	sumValue := sumFunction(inputNumber)
	fmt.Println("mod 3 + mod 5 sum: " + strconv.FormatInt(sumValue, 10))
}

func sumFunction(input int64) int64 {
	sum := int64(0)
	for input > 0 {
		if input%3 == 0 || input%5 == 0 {
			sum = sum + input
		}
		input = input - 1
	}
	return sum
}
