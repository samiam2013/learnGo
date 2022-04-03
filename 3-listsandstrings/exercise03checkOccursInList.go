package listsandstrings

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// Write a function that checks whether an element occurs in a list.
func exercise03() {
	// seed a random number generator with the time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("CQ CQ CQ")
	// create a search space for the function
	searchSpace := list.New()
	for i := 0; i <= 100; i++ {
		searchSpace.PushBack(r.Intn(1000))
	}
	// differ output based on finding element
	if search(searchSpace, 73) {
		fmt.Println("5 by 9 (strong signal)")
	} else {
		fmt.Println("QRO (increase power)")
	}
}

func search(l *list.List, value interface{}) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return true
		}
	}
	return false
}
