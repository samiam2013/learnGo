package main

import "fmt"

// Write a function that tests whether a string is a palindrome.
func exercise06() {
	palindrome := "racecar"
	nonPalindrome := "nascar"
	printCheck(palindrome)
	printCheck(nonPalindrome)
}

func checkPalindrome(candidate string) bool {
	startIndex := 0
	endIndex := len(candidate) - 1
	for i := 0; i <= len(candidate)/2; i++ {
		if candidate[startIndex+i] != candidate[endIndex-i] {
			return false
		}
	}
	return true
}

func printCheck(candidate string) {
	if checkPalindrome(candidate) {
		fmt.Println(candidate, "is a palindrome!")
	} else {
		fmt.Println(candidate, "is not a palindrome!")
	}
}
