package listsAndStrings

import (
	"fmt"
	"log"
	"math"
)

// Ex01 : redundantly exports LargestInt. sorry
func Ex01() {
	numbers := []int32{830, 1, 58, 282, 500, 32, 84}
	fmt.Println("largest number in list: ", LargestInt(numbers))
}

// LargestInt : 1 Write a function that returns the largest element in a list.
//  there are lists in golang but I'm going to use a slice.
func LargestInt(list []int32) int32 {
	// lol the compiler doesn't know this is an int32
	if len(list) == 0 || list == nil {
		log.Fatal("list for largestInt must be non-nil & populated.")
	}
	largest := int32(math.MinInt32)
	for i := 0; i < len(list); i++ {
		if list[i] > largest {
			largest = list[i]
		}
	}
	return largest
}
