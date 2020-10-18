package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Write a function that takes a list of strings an prints them, one per line,
//  in a rectangular frame. For example the list
//  ["Hello", "World", "in", "a", "frame"] gets printed as:
// *********
// * Hello *
// * World *
// * in    *
// * a     *
// * frame *
// *********
func main() {
	stringSlice := []string{"Hello", "World", "in", "a", "frame"}
	printFramed(stringSlice)
}

func printFramed(stringSlice []string) {
	maxLen := maxStrLen(stringSlice)

	for row := 0; row < len(stringSlice)+2; row++ {
		if row == 0 {
			fmt.Println(strings.Repeat("_", maxLen+4))
		} else if row == len(stringSlice)+1 {
			fmt.Println(strings.Repeat("‾", maxLen+4))
		} else {
			formatString := "%-" + strconv.Itoa(maxLen) + "v"
			fmt.Println("| " + fmt.Sprintf(formatString, stringSlice[row-1]) + " |")
		}
	}
}

func maxStrLen(stringSlice []string) int {
	length := 0
	for _, val := range stringSlice {
		if length < len(val) {
			length = len(val)
		}
	}
	return length
}
