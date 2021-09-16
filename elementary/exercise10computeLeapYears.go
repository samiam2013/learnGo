package elementary

import (
	"fmt"
	"time"
)

// Write a program that prints the next 20 leap years

// Next20LeapYears dynamically gets the current year and calculates
func Next20LeapYears() {
	leapYear, _, _ := time.Now().Date()
	leapYear += 4 - leapYear%4
	for i := 0; i < 20; i++ {
		if leapYear%100 == 0 && leapYear%400 != 0 {
			leapYear += 4
		}
		fmt.Println(leapYear, "is a leap year.")
		leapYear += 4
	}
}
