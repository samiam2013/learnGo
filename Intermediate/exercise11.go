package main

import "fmt"

/*
  If you chose your growth right in the previous problem, you typically won’t
    allocate very often. However, adding to a big list sometimes consumes
    considerable time. That might be problematic in some applications.
  Instead try allocating new chunks of memory for new items.
    So when your list is full and the user wants to add something,
    allocate a new chunk of 100 elements instead of copying all elements
    over to a new large chunk. Think about where to do the book-keeping about
    which chunks you have. Different book keeping strategies can quite
    dramatically change the performance characteristics of your list.
*/

func (l *List) growAppend() int {
	newLen := 100
	newSlice := make([]*Element, newLen)
	l.listSlice = append(l.listSlice, newSlice...)
	return newLen
}

func exercise11() {
	fmt.Println("exercise 11...I'm going to cheat and change the algorithm to append")
	myList := NewList(100)
	fmt.Println("length of list before additions: ", len(myList.listSlice))
	fmt.Println("pushing too many elements onto the list..")
	// hahaha just a few
	for i := 0; i < 100000025; i++ {
		myList.PushBack(i)
	}
	fmt.Println("length of list after elements added", len(myList.listSlice))
}
