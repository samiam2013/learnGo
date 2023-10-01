package elementary

import (
	"math/big"
)

// Write a program that computes the sum of an alternating series
//  where each element of the series is an expression of the form
//  ((-1)^{k+1})/(2 * k-1)
// for each value of k from 1 to a million, multiplied by 4.

// AlternatingSeries computes the series in exercise 11 and prints the reuslt
func AlternatingSeriesBig() *big.Float {
	i := big.NewInt(1)
	limit := big.NewInt(1000000)
	sum := big.NewFloat(0)
	for i.Cmp(limit) == -1 {
		k := *i
		k.Add(&k, big.NewInt(1))
		numerator := big.NewInt(0)
		numerator.Exp(big.NewInt(-1), &k, nil)

		k.Sub(&k, big.NewInt(1)) // k = k - 1 because 1 was added above and it's reused

		denominator := big.NewInt(0)
		denominator.Mul(big.NewInt(2), &k)
		denominator.Sub(denominator, big.NewInt(1))

		value := big.NewFloat(0)
		value.SetPrec(100000)
		value.SetInt(numerator)
		value.Quo(value, big.NewFloat(0).SetInt(denominator))

		sum.Add(sum, value)
		// fmt.Println("sum, value: ", sum, value)
		i.Add(i, big.NewInt(1))
	}
	return sum.Mul(sum, big.NewFloat(4))
}
