package main

import (
	"fmt"
)

type myElement struct {
	Next, Prev *myElement
	list       *myList
	Value      interface{}
}
type myList struct {
	root *myElement // sentinel list element, only &root, root.prev, and root.next are used
	len  int        // current list length excluding (this) sentinel element
}

// Write function that reverses a list, preferably in place.
func exercise02() {
	root := myElement{nil, nil, nil, "root"}
	strings := myList{&root, 1}
	strSlice := []string{"These", "values", "to", "be", "loaded", "into", "list"}
	for i := 0; i < len(strSlice); i++ {

	}
	fmt.Println("list forwards: ")
	printMyList(&strings)
	reverseList(&strings)

}

func printMyList(l *myList) {
	for e := l.root; e != nil; e = e.Next {
		fmt.Println(e.Value)
	}
}

// I'll be damned if I can't reverse this in-place with reflection.
func reverseList(l *myList) {
	for e := l.root; e != nil; e = e.Next {
	}
}
