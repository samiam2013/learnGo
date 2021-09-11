package elementary

import (
	"fmt"
	"math/big"
)

// Write a program that prints all prime numbers.
//  (Note: if your programming language
//  does not support arbitrary size numbers,
//  printing all primes up to the largest number
//  you can easily represent is fine too.)

// PrimeSeive prints an infinite number of primes
// 	takes an int64 that represents a limit, -1 for no limit
//  takes a bool value to either print or simply run
func PrintPrimeSeive(nToFind int64, print bool) {
	// start with the smallest prime, 3
	candidate := new(big.Int)
	candidate.SetInt64(3)
	for i := new(big.Int).SetInt64(0); nToFind == -1 || i.Cmp(new(big.Int).SetInt64(nToFind)) <= 0; {
		candidate.Add(candidate, new(big.Int).SetInt64(2))
		if CheckBigPrime(candidate) {
			if print {
				fmt.Println(candidate, " is prime!")
			}
			// increment number of primes found by 1
			i.Add(i, new(big.Int).SetInt64(1))
		}
	}
}

func CheckBigPrime(candidate *big.Int) bool {
	// create a max denominator by taking the square root and adding 1
	maxDenominator := new(big.Int)
	maxDenominator.Sqrt(candidate)
	maxDenominator.Add(maxDenominator, new(big.Int).SetInt64(1))
	divisor := new(big.Int).SetInt64(2)
	modulus := new(big.Int).SetInt64(1)
	// from 2 -> max denominator try modulus-ing the prime and check == 0
	for maxDenominator.Cmp(divisor) >= 0 {
		modulus.Mod(candidate, divisor)
		if modulus.Cmp(new(big.Int).SetInt64(0)) == 0 {
			return false
		}
		divisor.Add(divisor, new(big.Int).SetInt64(1))
	}
	return true
}
