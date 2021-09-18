package intermediate

import "fmt"

// ListI interface enforces basic functionality for type List
type ListI interface {
	Len() int
	Front() *Element
	Back() *Element
	PushBack(interface{}) *Element
	PushFront(interface{}) *Element
	Remove(*Element) interface{}
}

// ElementI interface enforces getters for hidden pointers
type ElementI interface {
	Next() *Element
	Prev() *Element
}

// Think of a good interface for a list. What operations do you typically need?
// You might want to investigate the list interface in your language and in some
//  other popular languages for inspiration.
func exercise08() {
	fmt.Println("this file just implements interfaces for the next exercise.")
}
