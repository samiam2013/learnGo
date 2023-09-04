package main

import (
	"fmt"
	"time"
)

var parity = 0

func main() {
	bits := 8 * 1024 // 1 KB
	for h := 0; h < 10; h++ {
		randBits := getRandBits(bits)
		timeStart := time.Now()
		for i := 0; i < 10; i++ {
			parity = getParity(randBits)
		}
		timeTaken := time.Since(timeStart)
		fmt.Printf("Time taken for %d kb: %v \ttime(ns) per kb: %v\n",
			(bits / (8 * 1024)), timeTaken, int(timeTaken.Nanoseconds())/(bits/(8*1024)))
		bits += 10000 * 8 * 1024 // 100 kb
	}
}

func getRandBits(bits int) []byte {
	randBits := make([]byte, bits)
	for i := 0; i < bits; i++ {
		randBits[i] = byte(i % 2)
	}
	return randBits
}

func getParity(randBits []byte) int {
	parity := 0
	for _, bit := range randBits {
		parity ^= int(bit)
	}
	return parity
}
