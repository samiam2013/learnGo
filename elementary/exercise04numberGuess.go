package main

import (
	"fmt"
	"log"
	"strconv"
)

func exercise04() {
	var inputNumberString string
	fmt.Print("Pick a number, an number that doesn't overflow, of course: ")
	fmt.Scanln(&inputNumberString)
	number, err := strconv.ParseInt(inputNumberString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inputNumberString + "! = " +
		strconv.FormatInt(triangular(number), 10))
}
