package main

import "fmt"
import "strconv"

// Write a function that takes a number and returns a list of its digits.
//  So for 2342 it should return [2,3,4,2].
func main() {
	slice := stringToIntSlice("7300")
	fmt.Println(slice)
}

func stringToIntSlice(input string) []int {
	slice := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		slice[i], err = strconv.Atoi(string(input[i]))
		if err != nil {
			return nil
		}
	}
	return slice
}
