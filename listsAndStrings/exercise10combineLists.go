package listsAndStrings

import (
	"container/list"
)

// Write a function that combines two lists by alternatingly taking elements,
//  e.g. [a,b,c], [1,2,3] → [a,1,b,2,c,3].
func exercise10() {
	l1 := list.New()
	l2 := list.New()
	for _, v := range []int8{1, 2, 3} {
		l1.PushBack(v)
	}
	for _, v := range []rune{'a', 'b', 'c'} {
		l2.PushBack(v)
	}
	l3 := interList(l2, l1)
	printList(l3)
}

func interList(l1, l2 *list.List) *list.List {
	sumList := list.New()
	e1 := l1.Front()
	e2 := l2.Front()
	var endOfValues bool
	for true {
		endOfValues = true
		if e1 != nil {
			sumList.PushBack(e1.Value)
			e1 = e1.Next()
			endOfValues = false
		}
		if e2 != nil {
			sumList.PushBack(e2.Value)
			e2 = e2.Next()
			endOfValues = false
		}
		if endOfValues {
			break
		}
	}
	return sumList
}
