package main

import (
	"container/list"
)

// Write a function that merges two sorted lists into a new sorted list.
//  [1,4,6],[2,3,5] → [1,2,3,4,5,6].
// You can do this quicker than concatenating them followed by a sort.
func exercise11() {
	l1 := list.New()
	for _, v := range []int8{1, 4, 6, 7, 8, 9} {
		l1.PushBack(v)
	}
	l2 := list.New()
	for _, v := range []int8{2, 3, 5, 10, 11} {
		l2.PushBack(v)
	}
	printList(mergeSorted(l1, l2))
}

func mergeSorted(l1, l2 *list.List) *list.List {
	e1 := l1.Front()
	e2 := l2.Front()
	mergedList := list.New()
	for true {
		if e1 != nil && e2 != nil {
			if e1.Value.(int8) < e2.Value.(int8) {
				mergedList.PushBack(e1.Value)
				e1 = e1.Next()
			} else {
				mergedList.PushBack(e2.Value)
				e2 = e2.Next()
			}
		} else if e1 != nil {
			mergedList.PushBack(e1.Value)
			e1 = e1.Next()
		} else if e2 != nil {
			mergedList.PushBack(e2.Value)
			e2 = e2.Next()
		} else {
			break
		}
	}
	return mergedList
}
