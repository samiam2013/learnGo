package elementary

import "fmt"

// Write a program that prints the next 20 leap years
func exercise10() {
	leapYear := 2024
	for i := 0; i < 20; i++ {
		if leapYear%100 == 0 && leapYear%400 != 0 {
			leapYear += 4
		}
		fmt.Println(leapYear, "is a leap year.")
		leapYear += 4
	}
}
