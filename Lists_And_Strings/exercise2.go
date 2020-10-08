package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// Write function that reverses a list, preferably in place.
func main() {
	strings := list.New()
	strSlice := []string{"These", "values", "to", "be", "loaded", "into", "list"}
	for i := 0; i < len(strSlice); i++ {
		strings.PushBack(strSlice[i])
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

func reverseList(l *list.List) {
	var elementBuffer *list.Element
	for e := l.Front(); e != nil; e = e.Next() {
		// first element previous will be empty
		elementBuffer = reflect.ValueOf(e).Interface().(*list.Element)
		elementBuffer.next
		//elementBuffer = e.next
		//e.next = e.prev
		//e.prev = elementBuffer
	}
}
