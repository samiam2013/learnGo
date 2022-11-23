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
	//lastFreq := 2
	for i := float64(2); i <= 1000000; i++ {
		if extraDay > (1 / i) {
			fmt.Println("have to have a leap year every " + strconv.Itoa(int(i-1)) + " years")
			extraDay -= (1 / (i - 1))
			//lastFreq = int(i - 1)
		} else if math.Abs(extraDay) > (1/i) && int(i)%100 == 0 {
			fmt.Println("have to remove a leap year every " + strconv.Itoa(int(i)) + " years")
			extraDay += (1 / i)
		}
		//fmt.Printf("extra day: %f, 1/i: %f\n", extraDay, 1/i)
	}
	fmt.Println("remaining fraction day per year: ", extraDay)
}
