package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Implement binary search.
func main() {
	slice := makeSortedIntSlice(10000)
	searchVal := 1292
	if binarySearch(slice, searchVal) {
		fmt.Println("found it!")
	} else {
		fmt.Println("couldn't find search value")
	}
}

func binarySearch(slice []int, searchVal int) bool {
	median := len(slice) / 2
	jumpDist := len(slice) / 4
	tinyJumps := 0
	for true {
		if slice[median] > searchVal {
			median -= jumpDist
		} else if slice[median] < searchVal {
			median += jumpDist
		} else if slice[median] == searchVal {
			return true
		}
		jumpDist = (jumpDist / 2)
		if jumpDist < 1 {
			jumpDist = 1
			tinyJumps++
			if tinyJumps > 2 {
				break
			}
		}
	}
	return false
}

func makeSortedIntSlice(length int) []int {
	seed := rand.NewSource(time.Now().UnixNano())
	randomer := rand.New(seed)
	slice := make([]int, length)
	slice[0] = 1
	for i := 1; i < length; i++ {
		slice[i] = slice[i-1] + randomer.Intn(3)
	}
	return slice
}
