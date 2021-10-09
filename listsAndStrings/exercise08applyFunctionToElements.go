package listsAndStrings

import (
	"container/list"
	"math"
)

// Write a function on_all that applies a function to every element of a list.
// Use it to print the first twenty perfect squares. The perfect squares can be
//  found by multiplying each natural number with itself.
// The first few perfect squares are 1*1= 1, 2*2=4, 3*3=9, 4*4=16.
// Twelve for example is not a perfect square because there is no natural number
//  m so that m*m=12.
func exercise08() {
	perfectSquares := list.New()
	for i := 1; i <= 20; i++ {
		perfectSquares.PushBack(int64(i))
	}
	squareValue := func(e *list.Element) {
		e.Value = int64(math.Pow(float64(e.Value.(int64)), float64(2)))
	}
	onAll(squareValue, perfectSquares)
	printList(perfectSquares)
}

func onAll(runFunc func(e *list.Element), l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		runFunc(e)
	}
}
