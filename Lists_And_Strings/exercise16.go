package main

import (
	"fmt"
	"math"
)

// Write a function that takes a list of numbers, a starting base b1
//  and a target base b2 and interprets the list as a number in base b1
//  and converts it into a number in base b2 (in the form of a list-of-digits).
// So for example [2,1,0] in base 3 gets converted to base 10 as [2,1].
func main() {
	base3Slice := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	base10Slice := convertSliceBase(base3Slice, 2, 10)
	fmt.Println(base10Slice)
}

func convertSliceBase(baseXSlice []int, x, y int) []int {
	place := 0
	sum := 0
	for i := len(baseXSlice) - 1; i >= 0; i-- {
		placePower := math.Pow(float64(x), float64(place))
		//fmt.Println(placePower)
		sum += baseXSlice[i] * int(placePower)
		place++
	}
	newLength := int(math.Floor(math.Log(float64(sum))/math.Log(float64(y)))) + 1
	fmt.Println("new length ", newLength)
	fmt.Println("sum value ", sum)
	baseYSlice := make([]int, newLength)
	place = 0
	for i := newLength - 1; i >= 0; i-- {
		placePower := math.Pow(float64(y), float64(place))
		//fmt.Println(placePower)
		remainder := (sum / int(placePower)) % 10
		sum -= remainder
		baseYSlice[i] = remainder
		place++
	}
	return baseYSlice
}
