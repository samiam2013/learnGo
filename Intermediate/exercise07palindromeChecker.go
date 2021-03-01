package main

import "fmt"

// Write a program that finds the longest palindromic substring of a given string.
//  Try to be as efficient as possible!
func exercise07() {
	palindrome := findLongestPalindrome("Hello hannah, do you like racecars?")
	fmt.Println(palindrome)
}

func findLongestPalindrome(text string) string {
	fmt.Println("starting search")
	palindrome := ""
	j, k := 0, 0
	for i := 0; i < len(text)-2; i++ {
		if text[i] == text[i+1] {
			// get new iterators j, k and find bounds
			j, k = i, i+1
			for text[j] == text[k] {
				if j > 0 && k < len(text)-1 {
					j--
					k++
				}
			}
			j++
			k--
		} else if text[i] == text[i+2] {
			// get new iterators j, k and find bounds
			j, k = i, i+2
			for text[j] == text[k] {
				if j > 0 && k < len(text)-1 {
					j--
					k++
				}
			}
			j++
			k--
		} else {
			continue
		}
		if j != 0 || k != 0 && (k-j > len(palindrome)) {
			palindrome = text[j : k+1]
		}
	}
	return palindrome
}
