package elementary

import "math/big"

// Write a program that computes the sum of an alternating series
//  where each element of the series is an expression of the form
//  ((-1)^{k+1})/(2 * k-1)
// for each value of k from 1 to a million, multiplied by 4.

// AlternatingSeries computes the series in exercise 11 and prints the reuslt
func AlternatingSeries() *big.Float {
	// sum := float64(0)
	// for k := float64(1); k <= 1000000; k++ {
	// 	value := (math.Pow((-1), ((k * 4) + 1))) / (2*(k*4) - 1)
	// 	sum += value
	// 	//fmt.Println("sum: ", sum, value)
	// }
	// return sum

	sum := big.NewFloat(0)
	for i := big.NewFloat(0); i.Cmp(big.NewFloat(10000)) <= 0; i.Add(i, big.NewFloat(1)) {
		negOne := big.NewFloat(-1)
		two := big.NewFloat(2)
		one := big.NewFloat(1)
		four := big.NewFloat(4)
		k := big.NewFloat(0)
		k.Copy(i)
		numerator := Exp(negOne, (k.Mul(k, four)).Add(k, one))
		denominator := two.Mul(two, (k.Mul(k, four))).Sub(k, one)
		value := numerator.Quo(numerator, denominator)
		sum.Add(sum, value)
		i.Add(i, one)

	}
	return sum.Mul(sum, big.NewFloat(4))
}

func Exp(x, y *big.Float) (z *big.Float) {
	//  exponentiation by multiplication
	z = big.NewFloat(1)
	for y.Cmp(big.NewFloat(0)) > 0 {
		z.Mul(z, x)
		y.Sub(y, big.NewFloat(1)) // y--
	}
	return
}
