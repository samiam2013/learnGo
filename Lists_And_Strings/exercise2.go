package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type element struct {
	Next, Prev *element
	list       *list
	Value      interface{}
}
type list struct {
	root element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Write function that reverses a list, preferably in place.
func main() {
	root := element{nil, nil, nil, "root"}
	strings := list{root, 1}
	strSlice := []string{"These", "values", "to", "be", "loaded", "into", "list"}
	for i := 0; i < len(strSlice); i++ {

	}
	fmt.Println("list forwards: ")
	printList(strings)
	reverseList(strings)

}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// I'll be damned if I can't reverse this in-place with reflection.
func reverseList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		// first element previous will be empty
		re := reflect.ValueOf(&e).Elem()
		reF := re.FieldByName("next")
		// haha unsafe go brrrrrr
		ele := reflect.NewAt(reF.Type(), unsafe.Pointer(reF.UnsafeAddr())).Elem()
		fmt.Println(ele)
		//elementBuffer = e.next
		//e.next = e.prev
		//e.prev = elementBuffer
	}
}
