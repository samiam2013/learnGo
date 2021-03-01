package main

import (
	"container/list"
)

// Write a function that concatenates two lists [a,b,c], [1,2,3] → [a,b,c,1,2,3]
func exercise09() {
	l1 := list.New()
	l2 := list.New()
	for _, v := range []rune{'a', 'b', 'c'} {
		l1.PushBack(v)
	}
	for _, v := range []int8{1, 2, 3} {
		l2.PushBack(v)
	}
	l1.PushBackList(l2)
	printListReflect(l1)
}
