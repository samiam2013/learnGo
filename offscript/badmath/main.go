package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i, i+i)
		fmt.Printf("%d x %d = %d\n", i, i, i*i)
		fmt.Printf("%d ^ %d = %d\n", i, i,
			int64(math.Pow(float64(i), float64(i))))
	}
	nativeTime := time.Since(start)
	start = time.Now()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i, add(i, i))
		fmt.Printf("%d x %d = %d\n", i, i, multiply(i, i))
		fmt.Printf("%d ^ %d = %d\n", i, i, power(i, i))
	}
	loopTime := time.Since(start)
	fmt.Printf("took %.2fs\n", loopTime.Seconds())
	speedFactorF := float64(loopTime) / float64(nativeTime)
	factor := int64(speedFactorF)
	factor -= factor % 100
	fmt.Printf("native took %s or %d times faster\n", nativeTime, factor)
}

func add(l ...int) int {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return sum
}

func multiply(x, y int) int {
	sum := 0
	for i := 1; i <= y; i++ {
		sum = add(x, sum)
	}
	return sum
}

func power(x, y int) int {
	sum := 1
	for i := 0; i < y; i++ {
		sum = multiply(x, sum)
	}
	return sum
}
