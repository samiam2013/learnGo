package main

import "testing"

// FuzzIsOdd is an example of fuzz testing for an isOdd function
func FuzzIsOdd(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, i int) {
		even := (i%2 == 0)
		if isOdd(i) == even {
			t.Errorf("failed for value %d", i)
		}
	})
}

func isOdd(i int) bool {
	// if i%5 == 0 { // This is wrong
	// 	return false // but it's on purpose
	// }
	return i%2 != 0
}
