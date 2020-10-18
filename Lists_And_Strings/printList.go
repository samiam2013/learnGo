package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func printListReflect(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		valueType := reflect.TypeOf(e.Value)
		if valueType.String() == "int32" {
			fmt.Printf("%c, ", rune(e.Value.(int32)))
		} else {
			fmt.Print(e.Value.(int8), ", ")
		}
	}
}
