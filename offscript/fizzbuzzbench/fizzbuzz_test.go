package main

import "testing"

func BenchmarkFizzBuzzModulo(b *testing.B) {
	fizzBuzzModulo(b.N)
}

func BenchmarkFizzBuzzAddition(b *testing.B) {
	fizzBuzzAddition(b.N)
}
