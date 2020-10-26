package main

import "fmt"

// List consists of root node
type List struct {
	listSlice   []*Element
	first, last int
}

// Front implements requirment
func (l *List) Front() *Element {
	return l.listSlice[l.first]
}

// Back implements requirement
func (l *List) Back() *Element {
	return l.listSlice[l.last]
}

// Len implements requirment
func (l *List) Len() int {
	// at first == last == 0, 1 element is used.
	return l.last - l.first + 1
}

// PushBack implements requirment
func (l *List) PushBack(v interface{}) *Element {
	if l.last < len(l.listSlice)-1 {
		newElem := &Element{v, l.last + 1, l}
		if l.listSlice[l.last] == nil {
			l.last--
		}
		// there's still space!
		l.listSlice[l.last+1] = newElem
		l.last++
		return newElem
	}
	panic("ran out of length in list for PushBack()")
}

// PushFront implements requirment
func (l *List) PushFront(v interface{}) *Element {
	if l.last < len(l.listSlice)-1 {
		// TODO FINISH
		var buf *Element
		var i int
		for i = l.last; i >= l.first; i-- {
			buf = l.listSlice[i+1]
			l.listSlice[i+1] = l.listSlice[i]
			l.listSlice[i] = buf
		}
		l.listSlice[l.first] = &Element{v, i + 1, l}
		l.last++
		// this is a hack and I don't know why I need it.
		if l.listSlice[l.last] == nil {
			l.last--
		}
		return l.listSlice[l.first]
	}
	panic("ran out of length in list for PushFront()")
}

// Remove implements requirement
func (l *List) Remove(e *Element) interface{} {
	// TODO FINISH
	var i int
	for i = e.index; i < l.last; i++ {
		l.listSlice[i] = l.listSlice[i+1]
	}
	l.last--
	return e.Value
}

// Element consists of a value, next and prev element pointers
type Element struct {
	Value interface{}
	index int
	list  *List
}

// Next fills requirment of ElementI
func (e *Element) Next() *Element {
	if e.index+1 <= e.list.last {
		return e.list.listSlice[e.index+1]
	}
	return nil
}

// Prev fills requirment of ElementI
func (e *Element) Prev() *Element {
	if e.index-1 >= e.list.first {
		return e.list.listSlice[e.index-1]
	}
	return nil
}

// Implement your list interface using a fixed chunk of memory, say an array of
//  size 100. If the user wants to add more stuff to your list than fits in your
//  memory you should produce some kind of error, for example you can throw an
//  exception if your language supports that.
func exercise09() {
	// my list implementation is very manual
	//  and I can't keep go from panicking if you use PushFront()
	//  for the first element, so I hacked it to check
	//   if l.listSlice[l.last] == nil and decrement if it is.
	var l ListI
	lSlice := make([]*Element, 100)
	l = &List{lSlice, 0, 0}

	l.PushFront("Hello")
	l.PushBack("World")

	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)
}
