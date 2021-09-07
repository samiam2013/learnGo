package elementary

import (
	"fmt"
	"log"
	"strconv"
)

// TriangleSum : 4
func TriangleSum() {
	var inputNumberString string
	fmt.Print("Pick a number, an number that doesn't overflow, of course: ")
	fmt.Scanln(&inputNumberString)
	number, err := strconv.ParseInt(inputNumberString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inputNumberString + "?: " +
		strconv.FormatInt(triangular(number), 10))
}

func triangular(input int64) int64 {
	sum := int64(0)
	for input > 0 {
		sum = sum + input
		input = input - 1
	}
	return sum
}
