package main

import (
	"testing"

	"golang.org/x/sys/unix"
)

func init() {
	var set unix.CPUSet
	_ = unix.SchedGetaffinity(0, &set)
	_ = unix.SchedSetaffinity(0, &set)
}

func BenchmarkFizzBuzzModulo(b *testing.B) {
	fizzBuzzModulo(b.N)
}

func BenchmarkFizzBuzzAddition(b *testing.B) {
	fizzBuzzAddition(b.N)
}
