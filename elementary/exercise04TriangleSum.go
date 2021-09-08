package elementary

import (
	"fmt"
	"log"
	"strconv"
)

// TriangleSum : 4. Write a program that asks the user for a number n and prints the sum of the numbers 1 to n
func TriangleSum() {
	var inputNumberString string
	fmt.Print("Pick a number, an number that doesn't overflow, of course: ")
	fmt.Scanln(&inputNumberString)
	number, err := strconv.ParseInt(inputNumberString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inputNumberString + "?: " +
		strconv.FormatInt(Σ(number), 10))
}

func Σ(n int64) int64 {
	return (n * (n + 1)) / 2
}
