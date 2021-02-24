package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	if got != "fizzbuzz" {
		t.Errorf("FizzBuzz(15) = %s; want fizzbuzz", got)
	}
}
