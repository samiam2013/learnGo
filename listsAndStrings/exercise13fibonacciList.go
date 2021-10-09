package listsandstrings

import (
	"fmt"
	"math/big"
)

// Write a function that computes the list of the first 100 Fibonacci numbers.
// The first two Fibonacci numbers are 1 and 1. The n+1-st Fibonacci number
// can be computed by adding the n-th and the n-1-th Fibonacci number.
// The first few are therefore 1, 1, 1+1=2, 1+2=3, 2+3=5, 3+5=8
func exercise13() {
	n := new(big.Int).SetInt64(1)
	n1 := new(big.Int).SetInt64(1)
	sum := new(big.Int).SetInt64(0)
	for i := 0; i < 100; i++ {
		sum.Add(n, n1)
		n.Set(n1)
		n1.Set(sum)
		fmt.Println(sum)
	}
}
