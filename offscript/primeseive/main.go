package main

import (
	"math/big"
	"time"
)

func main() {

	limit := big.NewInt(100_000_000)
	c := make(chan *big.Int)
	big2 := big.NewInt(2)
	go func() {
		for i := big.NewInt(3); i.Cmp(limit) < 0; i.Add(i, big2) {
			iCpy := big.NewInt(0).Set(i)
			go func(i *big.Int) {
				sieve(i, c)
			}(iCpy)
		}
	}()
	for {
		select {
		case i := <-c:
			println(i.String())

		case <-time.After(1 * time.Second):
			return

		}
	}
}

func sieve(i *big.Int, c chan *big.Int) {
	big2 := big.NewInt(2)
	// from 2 to the square root of i (inclusive) check if i is divisible by j
	sqrt := big.NewInt(0).Sqrt(i)
	sqrt = sqrt.Add(sqrt, big.NewInt(1))
	// fmt.Println("candidate: ", i.String(), " sqrt: ", sqrt.String())
	for j := big.NewInt(3); j.Cmp(sqrt) < 0; j.Add(j, big2) {
		if big.NewInt(0).Mod(i, j).Cmp(big.NewInt(0)) == 0 {
			return
		}
	}
	c <- i
}
