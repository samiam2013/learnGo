package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

// Write a guessing game where the user has to guess a secret number.
//  After every guess the program tells the user whether their number
//  was too large or too small. At the end the number of tries needed
//  should be printed. It counts only as one try if they input the
//  same number multiple times consecutively.
func main() {
	randNum := int64(rand.Intn(1000))
	var guesses []int64
	fmt.Print("What is your guess for the random number?: ")
	var input string
	fmt.Scanln(&input)
	guess, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	for guess != randNum {
		if guess < randNum {
			fmt.Print("Higher!: ")
		} else if guess > randNum {
			fmt.Print("Lower!: ")
		} else {
			break
		}
		fmt.Scanln(&input)
		guess, err = strconv.ParseInt(input, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		if !inList(guess, guesses) {
			guesses = append(guesses, guess)
		}
	}
	fmt.Println("your guess was correct! took", len(guesses), "tries.")
}

func inList(number int64, list []int64) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == number {
			return true
		}
	}
	return false
}
