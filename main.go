package main

import (
	"fmt"

	"github.com/samiam2013/learnGo/elementary"
	"github.com/samiam2013/learnGo/intermediate"
)

func main() {
	elementary.HelloWorld()
	fmt.Println("TriangleSum(10): ", elementary.Σ(10))
	fmt.Println("FizzBuzz sumFunction(10): ", elementary.SumFunction(10))
	fmt.Println("Factorial(10):", elementary.Factorial(10))
	elementary.MultiplicationTable()
	elementary.PrintPrimeSeive(10, true)
	elementary.PrintPrimeSeiveP(10)
	elementary.Next20LeapYears()
	// intermediate.Ex01()
	// listsandstrings.Ex01()
	intermediate.BinTreeHeapSort()
}
