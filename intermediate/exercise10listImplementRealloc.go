package intermediate

import "fmt"

/*
  Improve your previous implementation such that an arbitrary number of elements
   can be stored in your list. You can for example allocate bigger and bigger
   chunks of memory as your list grows, copy the old elements over and release
   the old storage. You should probably also release this memory eventually
   if your list shrinks enough not to need it anymore. //"should probably" -> no
  Think about how much bigger the new chunk of memory should be so that your
   performance won’t be killed by allocations.
  Increasing the size by 1 element for example is a bad idea.
*/

// NewList creates a new list with a given number of open elements
func NewList(elements int) *List {
	return &List{make([]*Element, elements), 0, 0}
}

// it would be trivial to reverse this operation and check when the usage
//  is low, but I'm kind of sick of this set of exercises.
func (l *List) grow() int {
	len := l.Len()
	newLen := len * 2
	newSlice := make([]*Element, newLen)
	copy(newSlice, l.listSlice)
	l.listSlice = newSlice
	return newLen
}

func exercise10() {
	myList := NewList(100)
	fmt.Println("length of list before additions: ", len(myList.listSlice))
	fmt.Println("pushing too many elements onto the list..")
	// hahaha just a few GBs
	for i := 0; i < 100000000; i++ {
		myList.PushBack(i)
	}
	fmt.Println("length of list after elements added", len(myList.listSlice))
}

/*// alternate method for when it was using panics, but this doesn't
//    call back into the same funtion and complete the push
func recoverByAllocate(l *List) {
	var newLen int
	if r := recover(); r != nil {
		newLen = l.grow()
	}
	fmt.Println("recovered panic, grew list to ", newLen, " elements")
}
*/
