package main

import (
	"container/list"
)

// Write a function that rotates a list by k elements. For example [1,2,3,4,5,6]
// rotated by two becomes [3,4,5,6,1,2]. Try solving this without creating a
// copy of the list. How many swap or move operations do you need?
func exercise12() {
	l := list.New()
	for _, v := range []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		l.PushBack(v)
	}
	rotateList(l, 3)
	printList(l)
}

func rotateList(l *list.List, rotations int) {
	for i := 0; i < rotations; i++ {
		l.PushBack(l.Remove(l.Front()))
	}
}
