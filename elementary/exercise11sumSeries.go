package elementary

import (
	"fmt"
	"math"
)

// Write a program that computes the sum of an alternating series
//  where each element of the series is an expression of the form
//  ((-1)^{k+1})/(2 * k-1)
// for each value of k from 1 to a million, multiplied by 4.
func AlternatingSeries() {
	sum := float64(0)
	for k := float64(1); k <= 1000000; k++ {
		value := (math.Pow((-1), ((k * 4) + 1))) / (2*(k*4) - 1)
		sum += value
		//fmt.Println("sum: ", sum, value)
	}
	fmt.Println("sum: ", sum)
}
