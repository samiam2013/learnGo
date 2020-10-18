package main

import (
	"fmt"
	"strconv"
)

// Write a program that outputs all possibilities to put + or - or nothing
// between the numbers 1,2,…,9 (in this order) such that the result is 100.
// For example 1  + 2 + 3 - 4 + 5 + 6 + 78 + 9 = 100.
func exercise01() {
	searchSpace(1, 2, 0, "0")
}

// depth first search of possibilities
func searchSpace(n, n1, sum int, expr string) {
	if n <= 9 {
		// add
		searchSpace(n1, n1+1, sum+n, expr+" + "+strconv.Itoa(n))
		// subtract
		searchSpace(n1, n1+1, sum-n, expr+" - "+strconv.Itoa(n))
		if n < 9 {
			// concatenated add
			searchSpace(n1+1, n1+2, sum+((n*10)+n1),
				expr+" + "+strconv.Itoa((n*10)+n1))
			// concatenated subtract
			searchSpace(n1+1, n1+2, sum-((n*10)+n1), expr+" - "+strconv.Itoa((n*10)+n1))
		}
	} else {
		if sum == 100 {
			fmt.Println(expr + " = 100")
		}
	}
}
