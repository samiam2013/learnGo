package main

import (
	"fmt"
	"math/big"
)

// Write a program that prints all prime numbers.
//  (Note: if your programming language
//  does not support arbitrary size numbers,
//  printing all primes up to the largest number
//  you can easily represent is fine too.)

func exercise08() {
	// start with the smallest prime, 3
	candidate := new(big.Int)
	candidate.SetInt64(3)
	maxDenominator := new(big.Int)
	for true {
		// create a max denominator by taking the square root and adding 1
		maxDenominator.Sqrt(candidate)
		maxDenominator.Add(maxDenominator, new(big.Int).SetInt64(1))
		divisor := new(big.Int).SetInt64(2)
		modulus := new(big.Int).SetInt64(1)
		isPrime := true
		// from 2 -> max denominator try modulus-ing the prime and check == 0
		for maxDenominator.Cmp(divisor) >= 0 {
			modulus.Mod(candidate, divisor)
			if modulus.Cmp(new(big.Int).SetInt64(0)) == 0 {
				isPrime = false
				break
			}
			divisor.Add(divisor, new(big.Int).SetInt64(1))
		}
		if isPrime {
			fmt.Println(candidate, " is prime!")
		}
		candidate.Add(candidate, new(big.Int).SetInt64(2))
	}
}
