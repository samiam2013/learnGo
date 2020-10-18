package main

import (
	"container/list"
	"fmt"
)

var runningTotal int64

// Write a function that computes the running total of a list.
func exercise05() {
	numbers := list.New()
	for _, v := range []int64{830, 73820, 2838, 83, 17, 1, 82038, 239329} {
		pushBackAndTotal(v, numbers)
	}
}

func pushBackAndTotal(value int64, l *list.List) {
	runningTotal += value
	l.PushBack(value)
	fmt.Println("running total of list: ", runningTotal)
	printList(l)
}
