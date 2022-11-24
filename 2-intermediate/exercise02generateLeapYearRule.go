package intermediate

import (
	"fmt"
	"math"
	"strconv"
)

// Write a program that takes the duration of a year (in fractional days) for an
//
//	imaginary planet as an input and produces a leap-year rule that minimizes
//	the difference to the planet’s solar year.
func exercise02() {
	fmt.Println("running exercise02()")
	GenerateLeapYearRule(365.2422)
}

func GenerateLeapYearRule(days float64) {
	extraDay := days - math.Floor(days)
	lastFreq := 2
	for i := float64(2); i <= 100_000; i++ {
		// skip if it's not a multiple of the last frequency
		// . or if it's less than 100 or not a multiple of 100
		if int(i)%lastFreq != 0 || !(i < 100 || int(i)%100 == 0) {
			continue
		}
		if extraDay >= ((1 / i) - ((1 / i) * 0.10)) {
			fmt.Println("have to have a leap year every " + strconv.Itoa(int(i)) + " years")
			extraDay -= (1 / (i))
			lastFreq = int(i)
		}
		if extraDay <= -((1 / i) + ((1 / i) * 0.10)) {
			fmt.Println("have to skip a leap year every " + strconv.Itoa(int(i)) + " years")
			extraDay += (1 / (i))
			lastFreq = int(i)
		}
	}
	fmt.Printf("remaining fraction day per year: %f\n", extraDay)
}
