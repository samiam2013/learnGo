package main

import "fmt"

// this is an alternative I found to fizzbuzz online
//  trying to make it to those 30 lines of go today.
func testAnagram() {
	fmt.Println("listen, silent: ", isAnagram("listen", "silent"))
	fmt.Println("sam, mass: ", isAnagram("sam", "mass"))
}

func isAnagram(word1, word2 string) bool {
	byte1 := []byte(word1)
	byte2 := []byte(word2)
	byte1 = bubbleSort(byte1)
	byte2 = bubbleSort(byte2)
	return string(byte1) == string(byte2)
}

func bubbleSort(bytes []byte) []byte {
	for i := 0; i < len(bytes); i++ {
		for j := 0; j < i; j++ {
			if bytes[j] > bytes[i] {
				buf := bytes[i]
				bytes[i] = bytes[j]
				bytes[j] = buf
			}
		}
	}
	return bytes
}
