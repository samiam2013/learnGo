package main

import "fmt"

// List interface enforces basic functionality for type List
type List interface {
	New() *List
	Len() int
	Front() *Element
	Back() *Element
	PushBack(interface{}) *Element
	PushFront(interface{}) *Element
	Remove(*Element) interface{}
}

// Element interface enforces getters for hidden pointers
type Element interface {
	Next() *Element
	Prev() *Element
}

// Think of a good interface for a list. What operations do you typically need?
// You might want to investigate the list interface in your language and in some
//  other popular languages for inspiration.
func exercise08() {
	fmt.Println("this file just implements interfaces for the next exercise.")
}
