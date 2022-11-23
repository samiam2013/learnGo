package intermediate

import (
	"fmt"
	"math"
	"strconv"
)

// Write a program that takes the duration of a year (in fractional days) for an
//  imaginary planet as an input and produces a leap-year rule that minimizes
//  the difference to the planet’s solar year.
func exercise02() {
	fmt.Println("running exercise02()")
	GenerateLeapYearRule(365.2422)
}

func GenerateLeapYearRule(days float64) {
	extraDay := days - math.Floor(days)
	for i := float64(2); i <= 10000; i++ {
		if extraDay > (1 / i) && extraDay < (1 / (i-1)) {
			fmt.Println("have to have a leap year every " + strconv.Itoa(int(i-1)) + " years")
			extraDay -= (1 / (i-1))
		}
	}
	fmt.Println("remaining fraction day per year: ", extraDay)
}
