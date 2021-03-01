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
	strings := myList{&root, 0}
	strSlice := []string{"These", "values", "to", "be", "loaded", "into", "list"}
	for _, val := range strSlice {
		strings.PushBack(val)
	}
	fmt.Println("list len: ", strings.len)
	fmt.Println("list forwards: ")
	strings.PrintMyList()
	reverseList(&strings)
	fmt.Println("list backwards: ")
	strings.PrintMyList()
}

func reverseList(l *myList) {
	var eBuf *myElement
	// e = e.Prev because increment operation is completed after
	//  pointers are reversed inside the loop
	for e := l.root.Next; e != nil && e.Value != "root"; e = e.Prev {
		// point Prev at Next
		eBuf = e.Prev
		e.Prev = e.Next
		// and Next at Prev through buffer
		e.Next = eBuf
	}
	// do the same for the root node
	eBuf = l.root.Prev
	l.root.Prev = l.root.Next
	l.root.Next = eBuf
}

func (l *myList) PushBack(value interface{}) {
	newElement := myElement{l.root, nil, l, value}
	if l.len == 0 {
		newElement.Prev = l.root
		l.root.Next = &newElement
	} else {
		newElement.Prev = l.root.Prev
		l.root.Prev.Next = &newElement
	}
	l.root.Prev = &newElement
	l.len++
}

func (l myList) PrintMyList() {
	for e := l.root.Next; e != nil && e.Value != "root"; e = e.Next {
		fmt.Println("element: ", e.Value)
	}
}
