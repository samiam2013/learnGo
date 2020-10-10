package main

import (
	"container/list"
	"fmt"
)

// Write a function that returns the elements on
//  odd positions in a list.
func main() {
	l := list.New()
	fmt.Println("the odd values of list: ")
	lSlice := []string{"I", "looked", "for", "someone", "to", "help"}
	for i := 0; i < len(lSlice); i++ {
		l.PushBack(lSlice[i])
	}
	oddList := extractOdd(l)
	printList(oddList)
}

func extractOdd(l *list.List) *list.List {
	returnList := list.New()
	counter := 1
	for e := l.Front(); e != nil; e = e.Next() {
		if counter%2 != 0 {
			returnList.PushBack(e.Value)
		}
		counter++
	}
	return returnList
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
