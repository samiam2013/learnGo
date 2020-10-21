package main

import "fmt"

// Write a program that finds the longest palindromic substring of a given string.
//  Try to be as efficient as possible!
func exercise07() {
	palindrome := findLongestPalindrome("Hello hannah, do you like racecars?")
	fmt.Println(palindrome)
}

func findLongestPalindrome(text string) string {
	palindrome := ""
	j, k := 0, 0
	for i := 1; i < len(text); i++ {
		if text[i] == text[i-1] {
			// get new iterators j, k and find bounds
			j, k = i, i-1
			for text[j] == text[k] {
				if j > 0 && k < len(text)-1 {
					j--
					k++
				}
			}
			j++
			k--
		} else if i > 1 && text[i] == text[i-2] {
			// get new iterators j, k and find bounds
			j, k = i, i-2
			for text[j] == text[k] {
				if j > 0 && k < len(text)-1 {
					j--
					k++
				}
			}
			j++
			k--
		}
		if j != 0 || k != 0 && (k-j > len(palindrome)) {
			palindrome = text[j : k+1]
		}
	}
	return palindrome
}
