package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// Write a function that concatenates two lists [a,b,c], [1,2,3] → [a,b,c,1,2,3]
func main() {
	l1 := list.New()
	l2 := list.New()
	for _, v := range []rune{'a', 'b', 'c'} {
		l1.PushBack(v)
	}
	for _, v := range []int8{1, 2, 3} {
		l2.PushBack(v)
	}
	l1.PushBackList(l2)
	printList(l1)
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		valueType := reflect.TypeOf(e.Value)
		if valueType.String() == "int32" {
			fmt.Printf("%c, ", rune(e.Value.(int32)))
		} else {
			fmt.Print(e.Value.(int8), ", ")
		}
	}
}
