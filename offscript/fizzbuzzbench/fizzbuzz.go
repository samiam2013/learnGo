package main

import (
	"strconv"
)

func fizzBuzzModulo(max int) {
	for i := 1; i <= max; i++ {
		value := ""
		if i%3 == 0 {
			value += "Fizz"
		}
		if i%5 == 0 {
			value += "Buzz"
		}
		if value == "" {
			value = strconv.Itoa(i)
		}
		_ = value
	}
}

func fizzBuzzAddition(max int) {
	nextFizz := 3
	nextBuzz := 5
	for i := 1; i <= max; i++ {
		value := ""
		if i == nextFizz {
			value += "Fizz"
			nextFizz += 3
		}
		if i == nextBuzz {
			value += "Buzz"
			nextBuzz += 5
		}
		if value == "" {
			value = strconv.Itoa(i)
		}
		_ = value
	}
}

func main() {

	// use `go test -bench=.`

	// const listLength = 100
	// results := [][]string{
	// 	fizzBuzzAddition(listLength),
	// 	fizzBuzzModulo(listLength),
	// }
	// for i := 0; i < len(results[0]); i++ {
	// 	if results[0][i] != results[1][i] {
	// 		log.Fatal("Error (mismatch): ", results[0][i], results[1][i])
	// 	}
	// }
}
